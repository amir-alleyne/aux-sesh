package services

import (
	"fmt"
	"net/http"
	"time"

	"github.com/amir-alleyne/aux-sesh/backend/models"
	"github.com/labstack/echo/v4"
)

/* Assume locked is aquired */
func CreateSession(c echo.Context, spotifyAdmin *models.SpotifyUser, globalSessions map[int]*models.Session) error {
	currentUser, err := GetUser(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	fmt.Println("Creating session with admin:", spotifyAdmin.ID)
	// check if the user is already in a session
	for _, session := range globalSessions {
		if isUserInSession(currentUser.ID, session) {
			return c.JSON(http.StatusConflict, "User is already in a session")
		}
	}
	if len(globalSessions) > 10 {
		return c.JSON(http.StatusTooManyRequests, "Too many sessions")
	}
	time := int(time.Now().Unix())

	if _, ok := globalSessions[time]; ok {
		return c.JSON(http.StatusInternalServerError, "Session with the same ID already exists")
	}
	session := models.Session{
		ID:      time,
		Admin:   spotifyAdmin,
		UserIDs: []string{spotifyAdmin.ID},
	}
	fmt.Println("Session created with ID:", session.ID)
	globalSessions[time] = &session
	return c.JSON(http.StatusOK, session)
}

func GetSessions(globalSessions map[int]*models.Session) []*models.Session {
	sessions := make([]*models.Session, 0)
	for _, session := range globalSessions {
		sessions = append(sessions, session)

	}
	return sessions
}

func isUserInSession(userID string, session *models.Session) bool {
	for _, id := range session.UserIDs {
		if id == userID {
			return true
		}
	}
	return false
}
