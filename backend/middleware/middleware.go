package middleware

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/zmb3/spotify"
)

// func SpotifyUserMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		c.Set("user", auth.AdminClient)
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
