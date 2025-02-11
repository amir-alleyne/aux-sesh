package sessions

import (
	"net/http"

	"github.com/labstack/echo/v4"
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
	return nil
}

func GetSessions(c echo.Context) error {
	// TODO : fix
	return nil
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
