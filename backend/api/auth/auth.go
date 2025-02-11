package auth

import (
	"fmt"
	"net/http"
	"os"
	"sync"

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
		htmlContent := fmt.Sprintf(`
          <html>
            <head>
              <script type="text/javascript">
                window.onload = function() {
                  window.opener.postMessage({
                    type: 'spotify-auth-callback',
                    isSignedIn: false,
                    error: %q
                  }, window.location.origin);
                  window.close();
                };
              </script>
            </head>
            <body>
              Authentication failed. Please close this window.
            </body>
          </html>
        `, errorMessage)
		return c.HTML(http.StatusForbidden, htmlContent)
	}

	client := Auth.NewClient(token)

	SessionsLock.Lock()
	defer SessionsLock.Unlock()
	spotifyUser := models.SpotifyUser{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		Expiry:       token.Expiry,
		Client:       &client,
	}
	services.CreateSession(c, &spotifyUser, Sessions)

	htmlContent := `
      <html>
        <head>
          <script type="text/javascript">
            window.onload = function() {
              window.opener.postMessage({
                type: 'spotify-auth-callback',
                isSignedIn: true
              }, window.location.origin);
              window.close();
            };
          </script>
        </head>
        <body>
          Successfully authenticated. Please close this window and return to the application.
        </body>
      </html>
    `
	return c.HTML(http.StatusOK, htmlContent)
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
