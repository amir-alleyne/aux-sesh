package services

import (
	"fmt"
	"sync"
	"time"

	"github.com/amir-alleyne/aux-sesh/backend/models"
	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/zmb3/spotify"
)

/* Assume locked is aquired */
func CreateSession(currentUser *clerk.User, spotifyAdmin *models.SpotifyUser, globalSessions map[int]*models.Session) (*models.Session, error) {
	// check if the user is already in a session
	for _, session := range globalSessions {
		if IsUserInSession(currentUser.ID, session) {
			return nil, fmt.Errorf("User is already in a session")
		}
	}
	if len(globalSessions) > 10 {
		return nil, fmt.Errorf("Too many sessions")
	}
	time := int(time.Now().Unix())

	if _, ok := globalSessions[time]; ok {
		return nil, fmt.Errorf("Session with the same ID already exists")
	}

	session := models.Session{
		ID:         time,
		Admin:      spotifyAdmin,
		AdminToken: spotifyAdmin.AccessToken,
		UserIDs:    []string{currentUser.ID},
		SongQueue:  []spotify.URI{},
		Lock:       sync.Mutex{},
	}
	fmt.Println("Session created with ID:", session.ID)
	globalSessions[time] = &session
	return &session, nil
}

func GetSessions(globalSessions map[int]*models.Session) []*models.Session {
	sessions := make([]*models.Session, 0)
	for _, session := range globalSessions {
		sessions = append(sessions, session)

	}
	return sessions
}

/*
Assume locked is aquired
*/
func JoinSession(user *clerk.User, globalSessions map[int]*models.Session, sessionID int) error {
	if globalSessions[sessionID] == nil {
		return fmt.Errorf("Session not found")
	}
	if IsUserInSession(user.ID, globalSessions[sessionID]) {
		return fmt.Errorf("User already in session")
	}
	globalSessions[sessionID].UserIDs = append(globalSessions[sessionID].UserIDs, user.ID)
	return nil
}

/*
Assume locked is aquired
*/
func LeaveSession(user *clerk.User, globalSessions map[int]*models.Session, sessionID int) error {
	for index, u := range globalSessions[sessionID].UserIDs {
		if user.ID == u {
			globalSessions[sessionID].UserIDs[index] = globalSessions[sessionID].UserIDs[len(globalSessions[sessionID].UserIDs)-1]
			globalSessions[sessionID].UserIDs = globalSessions[sessionID].UserIDs[:len(globalSessions[sessionID].UserIDs)-1]
			return nil
		}
	}
	return fmt.Errorf("User not in session")
}

/*
Assumes lock is acquired.
*/
func AddSongToQueue(session *models.Session, songID spotify.ID) error {
	err := session.Admin.Client.QueueSong(songID)
	return err
}

func IsUserInSession(userID string, session *models.Session) bool {
	for _, id := range session.UserIDs {
		if id == userID {
			return true
		}
	}
	return false
}
