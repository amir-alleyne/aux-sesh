package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/zmb3/spotify"

	"github.com/amir-alleyne/aux-sesh/backend/models"
	"github.com/labstack/echo/v4"
)

// func SpotifyUserMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		if auth.AdminClient != nil {
// 			c.Set("user", auth.AdminClient)
// 		}
// 		return next(c)
// 	}
// }

func GetUserFromContext(c echo.Context) (*spotify.PrivateUser, error) {
	// TODO : fix error when user is nil (occurs when user is not authenticated)
	userClient, ok := c.Get("user").(*spotify.PrivateUser)
	if !ok {
		return nil, fmt.Errorf("user not found in context")
	}
	fmt.Println("User found in context:", userClient.Email)
	return userClient, nil
}

func EnsureValidToken(user *models.SpotifyUser) (string, error) {
	if time.Now().After(user.Expiry) {
		err := RefreshToken(user)
		if err != nil {
			return "", fmt.Errorf("failed to refresh token: %v", err)
		}
	}
	return user.AccessToken, nil
}

func RefreshToken(user *models.SpotifyUser) error {
	data := url.Values{}
	clientID := os.Getenv("SPOTIFY_CLIENT_ID")
	clientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", user.RefreshToken)
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)

	req, _ := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	if newToken, ok := result["access_token"].(string); ok {
		user.AccessToken = newToken
		user.Expiry = time.Now().Add(time.Hour) // Update expiry
	}

	return nil
}
