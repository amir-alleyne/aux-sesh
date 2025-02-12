package services

import (
	"errors"
	"fmt"

	clrk "github.com/clerk/clerk-sdk-go/v2"

	"github.com/clerk/clerk-sdk-go/v2/jwt"
	"github.com/clerk/clerk-sdk-go/v2/user"
	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) (*clrk.User, error) {
	// Get the Authorization header (e.g., "Bearer <token>")
	token := c.Request().Header.Get("Authorization")
	//print headers

	if token == "" {
		if token == "" {
			// Attempt to retrieve token from cookie (if Clerk is set to store tokens there)
			cookie, err := c.Cookie("__session")

			if err != nil || cookie.Value == "" {
				return nil, errors.New("authorization token missing")
			}
			token = cookie.Value
		}
	} else {
		// Remove the "Bearer " prefix
		token = token[7:]

	}
	fmt.Println("token: ", token)
	// Verify the token with Clerk
	claims, err := jwt.Verify(c.Request().Context(), &jwt.VerifyParams{
		Token: token,
	})
	if err != nil {
		fmt.Println("error verifying token: %v", err)
		return nil, err
	}

	user, err := user.Get(c.Request().Context(), claims.Subject)
	if err != nil {
		return nil, err
	}

	return user, nil

}
