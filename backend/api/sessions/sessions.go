package sessions

import (
	"fmt"
	"net/http"
	"sync"
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

var (
	Sessions     = make(map[int]Session)
	SessionsLock sync.Mutex
)

/*
CreateSession is a handler function that creates a new session.
It should add the current user to the session and return a session ID.
*/
func CreateSession(c echo.Context) error {
	userClient, err := auth.GetAdmin()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	SessionsLock.Lock()
	defer SessionsLock.Unlock()

	if len(Sessions) > 10 {
		return c.JSON(http.StatusTooManyRequests, "Too many sessions")
	}
	time := unix.Now().Nanosecond()

	if _, ok := Sessions[time]; ok {
		return c.JSON(http.StatusInternalServerError, "Session with the same ID already exists")
	}
	session := Session{
		ID:      time,
		AdminID: userClient.Email,
		UserIDs: []string{userClient.Email},
	}
	fmt.Println("Session created with ID:", session.ID)
	Sessions[time] = session
	return c.JSON(http.StatusOK, session)
}

func EndSession(c echo.Context) error {
	return nil
}

func JoinSession(c echo.Context) error {
	return nil
}

func GetSessions(c echo.Context) error {
	sessions := make([]Session, 0)
	SessionsLock.Lock()
	defer SessionsLock.Unlock()
	for _, session := range Sessions {
		sessions = append(sessions, session)

	}
	return c.JSON(http.StatusOK, sessions)
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
