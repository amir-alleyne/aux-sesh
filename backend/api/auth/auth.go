package auth

import (
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/amir-alleyne/aux-sesh/backend/middleware"
	"github.com/amir-alleyne/aux-sesh/backend/models"
	"github.com/amir-alleyne/aux-sesh/backend/services"
	"github.com/clerkinc/clerk-sdk-go/clerk"
	log "github.com/sirupsen/logrus"

	"github.com/labstack/echo/v4"
	"github.com/zmb3/spotify"
)

var (
	// Change the redirect URI if needed.
	redirectURI = "http://localhost:8080/auth-callback"
	// Scopes required to control playback.
	Auth        = spotify.NewAuthenticator(redirectURI, spotify.ScopeUserReadPlaybackState, spotify.ScopeUserReadPrivate, spotify.ScopeUserReadEmail, spotify.ScopeUserModifyPlaybackState)
	State       = "some-random-string" // In production, generate a secure, random state value.
	ClerkClient *clerk.Client
	// In-memory storage for the admin’s Spotify client.
	AdminClient     *models.SpotifyUser
	AdminClientLock sync.RWMutex

	Sessions     = make(map[int]*models.Session)
	SessionsLock sync.Mutex
)

func SetAuth() (string, error) {
	clientID := os.Getenv("SPOTIFY_CLIENT_ID")
	clientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")

	if clientID == "" || clientSecret == "" {
		log.Fatal("Please set SPOTIFY_CLIENT_ID and SPOTIFY_CLIENT_SECRET environment variables.")
		return "", fmt.Errorf("Please set SPOTIFY_CLIENT_ID and SPOTIFY_CLIENT_SECRET environment variables.")
	}
	Auth.SetAuthInfo(clientID, clientSecret)
	c, err := (clerk.NewClient(os.Getenv("CLERK_SECRET_KEY")))
	if err != nil {
		log.Fatal("Error creating Clerk client:", err)
		return "", fmt.Errorf("Error creating Clerk client: %v", err)
	}

	ClerkClient = &c
	return clientID, nil
}

func Callback(c echo.Context) error {
	token, err := Auth.Token(State, c.Request())

	if err != nil {
		errorMessage := fmt.Sprintf("Couldn't get token: %v", err)
		return c.JSON(http.StatusForbidden, errorMessage)
	}

	client := Auth.NewClient(token)
	currentUser, err := services.GetUser(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	err = middleware.AddSpotifyTokenToMetaData(c, currentUser.ID, token.AccessToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	SessionsLock.Lock()
	defer SessionsLock.Unlock()
	spotifyUser := models.SpotifyUser{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		Expiry:       token.Expiry,
		Client:       &client,
	}

	session, err := services.CreateSession(currentUser, &spotifyUser, Sessions)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, session)
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
