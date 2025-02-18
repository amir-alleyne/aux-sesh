package sessions

import (
	"fmt"
	"net/http"

	"github.com/amir-alleyne/aux-sesh/backend/api/auth"
	"github.com/amir-alleyne/aux-sesh/backend/middleware"
	"github.com/amir-alleyne/aux-sesh/backend/models"
	"github.com/amir-alleyne/aux-sesh/backend/services"
	"github.com/labstack/echo/v4"
	"github.com/zmb3/spotify"
)

/*
CreateSession is a handler function that creates a new session.
It should add the current user to the session and return a session ID.

TODO: get the current user from the context and add them to the session.
*/
func CreateSession(c echo.Context) error {
	return nil
}

func EndSession(c echo.Context) error {
	return nil
}

func JoinSession(c echo.Context) error {
	currentUser, err := services.GetUser(c, true)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	var joinSessionRequest models.JoinSessionRequest
	err = c.Bind(&joinSessionRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if auth.Sessions[joinSessionRequest.SessionID] == nil {
		return c.JSON(http.StatusNotFound, "Session not found")
	}

	auth.SessionsLock.Lock()
	defer auth.SessionsLock.Unlock()
	session := auth.Sessions[joinSessionRequest.SessionID]
	err = middleware.ValidateAndUpdateToken(c, currentUser.ID, session.Admin)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, fmt.Sprintf("Admin refrsh token failed : %v", err))
	}

	err = services.JoinSession(currentUser, session)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "Joined session")
}

/*
	LeaveSession is a handler function that removes the current user from a session.

It should return an error if the user could not be removed.

x
*/
func LeaveSession(c echo.Context) error {
	currentUser, err := services.GetUser(c, true)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	var leaveSessionRequest models.JoinSessionRequest
	err = c.Bind(&leaveSessionRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if auth.Sessions[leaveSessionRequest.SessionID] == nil {
		return c.JSON(http.StatusNotFound, "Session not found")
	}

	auth.SessionsLock.Lock()
	defer auth.SessionsLock.Unlock()
	session := auth.Sessions[leaveSessionRequest.SessionID]

	err = services.LeaveSession(currentUser, auth.Sessions, session)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "Left session")
}

func GetSessions(c echo.Context) error {
	// TODO : fix
	return nil
}

/*
AddSongToQueue is a handler function that adds a song to the queue of the current session.
It should return an error if the song could not be added.
*/
func AddSongToQueue(c echo.Context) error {
	user, err := services.GetUser(c, true)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	var queueSongRequest models.QueueSongRequest
	if err := c.Bind(&queueSongRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	sessionID := queueSongRequest.SessionID
	songID := queueSongRequest.SongID

	auth.SessionsLock.Lock()
	session, exists := auth.Sessions[sessionID]
	auth.SessionsLock.Unlock()
	if !exists {
		return c.JSON(http.StatusNotFound, "Session not found")
	}

	if !services.IsUserInSession(user.ID, session) {
		return c.JSON(http.StatusForbidden, "User not in session")
	}
	err = middleware.ValidateAndUpdateToken(c, user.ID, session.Admin)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, fmt.Sprintf("Admin refrsh token failed : %v", err))
	}

	session.Lock.Lock()
	defer session.Lock.Unlock()

	services.AddSongToQueue(session, spotify.ID(songID))

	return c.JSON(http.StatusOK, "Song added to queue")
}

func PlaySong(c echo.Context) error {
	// songID := c.QueryParam("songID")
	// sessionID, err := strconv.Atoi(c.Param("session_id"))
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid session ID"})
	// }
	// userID := c.QueryParam("user_id")
	// session, exists := Sessions[sessionID]
	// if !exists {
	// 	return c.JSON(http.StatusNotFound, map[string]string{"error": "Session not found"})
	// }

	// if !isUserInSession(userID, session) {
	// 	return c.JSON(http.StatusForbidden, map[string]string{"error": "User not in session"})
	// }

	// _, err = middleware.EnsureValidToken(session.Admin)
	// if err != nil {
	// 	return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Admin authentication failed"})
	// }

	// session.Lock.Lock()
	// defer session.Lock.Unlock()

	// // play the song
	// err = session.Admin.Client.PlayOpt(&spotify.PlayOptions{
	// 	URIs: []spotify.URI{spotify.URI(songID)},
	// })
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, err.Error())
	// }
	return c.JSON(http.StatusOK, "Song playing")

}
