package sessions

import (
	"fmt"
	"net/http"
	unix "time"

	"github.com/amir-alleyne/aux-sesh/backend/api/auth"
	"github.com/labstack/echo/v4"
	"github.com/zmb3/spotify"
)

type Session struct {
	ID        int
	AdminID   string
	UserIDs   []string
	SongQueue []spotify.URI
}

/*
CreateSession is a handler function that creates a new session.
It should add the current user to the session and return a session ID.
*/
func CreateSession(c echo.Context) error {
	user := c.Get("user").(*spotify.PrivateUser)
	if user == nil {
		return c.JSON(http.StatusUnauthorized, "User not authenticated")
	}

	session := Session{
		ID:      unix.Now().Nanosecond(),
		AdminID: user.Email,
		UserIDs: []string{user.Email},
	}
	fmt.Println("Session created with ID:", session.ID)

	return c.JSON(http.StatusOK, "Session created")
}

func EndSession(c echo.Context) error {
	return nil
}

func JoinSession(c echo.Context) error {
	return nil
}

func PlaySong(c echo.Context) error {
	// songID := "spotify:track:6rqhFgbbKwnb9MLmUQDhG6"
	songID := c.QueryParam("songID")
	auth.AdminClientLock.RLock()
	client := auth.AdminClient
	defer auth.AdminClientLock.RUnlock()

	if client == nil {
		return c.JSON(http.StatusForbidden, "Admin not logged in")
	}

	// check if the song is in the queue

	// play the song
	err := client.PlayOpt(&spotify.PlayOptions{
		URIs: []spotify.URI{spotify.URI(songID)},
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "Song playing")

}
