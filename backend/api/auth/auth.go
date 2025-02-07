package auth

import (
	"fmt"
	"net/http"
	"os"
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/labstack/echo/v4"
	"github.com/zmb3/spotify"
)

var (
	// Change the redirect URI if needed.
	redirectURI = "http://localhost:8080/auth-callback"
	// Scopes required to control playback.
	Auth  = spotify.NewAuthenticator(redirectURI, spotify.ScopeUserReadPlaybackState, spotify.ScopeUserReadPrivate, spotify.ScopeUserReadEmail, spotify.ScopeUserModifyPlaybackState)
	State = "some-random-string" // In production, generate a secure, random state value.

	// In-memory storage for the admin’s Spotify client.
	AdminClient     *spotify.Client
	AdminClientLock sync.RWMutex
)

type User struct {
	ID    string
	Email string
}

func GetAdmin() (*spotify.PrivateUser, error) {
	AdminClientLock.RLock()
	defer AdminClientLock.RUnlock()

	if AdminClient == nil {
		return nil, fmt.Errorf("admin client not set")
	}
	spotifyUser, err := (*AdminClient).CurrentUser()
	if err != nil {
		return nil, fmt.Errorf("couldn't get user: %v", err)
	}
	return spotifyUser, nil
}

func SetAuth() (string, error) {
	clientID := os.Getenv("SPOTIFY_CLIENT_ID")
	clientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")

	if clientID == "" || clientSecret == "" {
		log.Fatal("Please set SPOTIFY_CLIENT_ID and SPOTIFY_CLIENT_SECRET environment variables.")
		return "", fmt.Errorf("Please set SPOTIFY_CLIENT_ID and SPOTIFY_CLIENT_SECRET environment variables.")
	}
	Auth.SetAuthInfo(clientID, clientSecret)

	return clientID, nil
}

func Callback(c echo.Context) error {
	token, err := Auth.Token(State, c.Request())
	if err != nil {
		errorMessage := fmt.Sprintf("Couldn't get token: %v", err)
		return c.JSON(http.StatusForbidden, errorMessage)
	}

	client := Auth.NewClient(token)

	AdminClientLock.Lock()
	AdminClient = &client
	AdminClientLock.Unlock()

	return c.JSON(http.StatusOK, "Admin login completed! You can now use playback features.")
}

// run server
func ClerkSignIn(c echo.Context) error {
	url := Auth.AuthURL(State)
	c.Redirect(http.StatusFound, url)
	return nil
}

func SpotifySignIn(c echo.Context) error {
	url := Auth.AuthURL(State)
	c.Redirect(http.StatusFound, url)
	return nil
}

func SignOut(c echo.Context) error {
	return nil
}
