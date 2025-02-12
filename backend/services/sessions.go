package services

import (
	"fmt"
	"sync"
	"time"

	"github.com/amir-alleyne/aux-sesh/backend/models"
	"github.com/clerk/clerk-sdk-go/v2"
)

/* Assume locked is aquired */
func CreateSession(currentUser *clerk.User, spotifyAdmin *models.SpotifyUser, globalSessions map[int]*models.Session) (*models.Session, error) {

	fmt.Println("Creating session with admin:", spotifyAdmin.ID)
	// check if the user is already in a session
	for _, session := range globalSessions {
		if isUserInSession(currentUser.ID, session) {
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
		ID:      time,
		Admin:   spotifyAdmin,
		UserIDs: []string{spotifyAdmin.ID},
		Lock:    sync.Mutex{},
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

func JoinSession() {

}
func isUserInSession(userID string, session *models.Session) bool {
	for _, id := range session.UserIDs {
		if id == userID {
			return true
		}
	}
	return false
}
