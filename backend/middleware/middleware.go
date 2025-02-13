package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	clerku "github.com/clerk/clerk-sdk-go/v2/user"

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

func ValidateAndUpdateToken(c echo.Context, clerkID string, user *models.SpotifyUser) error {
	accessToken, refreshed, err := EnsureValidToken(user)
	if err != nil {
		return fmt.Errorf("failed to validate token: %v", err)
	}
	if refreshed {
		err = AddSpotifyTokenToMetaData(c, clerkID, accessToken)
		if err != nil {
			return fmt.Errorf("failed to update token in metadata: %v", err)
		}
	}
	return nil
}

func EnsureValidToken(user *models.SpotifyUser) (string, bool, error) {
	var refreshed = false
	if time.Now().After(user.Expiry) {
		err := RefreshToken(user)
		if err != nil {
			return "", false, fmt.Errorf("failed to refresh token: %v", err)
		}
		refreshed = true
	}
	return user.AccessToken, refreshed, nil
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

	// TODO : add token to meta data

	return nil
}

func AddSpotifyTokenToMetaData(ctx echo.Context, userID, spotifyToken string) error {
	// Define the update parameters with the new public metadata.
	publicMetadata := map[string]interface{}{
		"spotify_token": spotifyToken,
	}

	// Marshal the map to JSON
	publicMetadataJSON, err := json.Marshal(publicMetadata)
	if err != nil {
		return err
	}

	// Convert JSON to json.RawMessage
	rawMessage := json.RawMessage(publicMetadataJSON)

	// Create the UpdateMetadataParams with the RawMessage
	metaData := &clerku.UpdateParams{
		PublicMetadata: &rawMessage,
	}

	// Call Clerk's Update function to update the user's metadata.
	updatedUser, err := clerku.Update(ctx.Request().Context(), userID, metaData)
	if err != nil {
		return fmt.Errorf("failed to update user metadata: %w", err)
	}

	fmt.Printf("Updated user: %v\n", updatedUser)
	return nil
}
