package middleware

import (
	"fmt"
	"net/http"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/zmb3/spotify"

	"github.com/amir-alleyne/aux-sesh/backend/api/auth"
	"github.com/amir-alleyne/aux-sesh/backend/api/sessions"
	"github.com/labstack/echo/v4"
)

func protectedWithAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, ok := clerk.SessionClaimsFromContext(r.Context())
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"access": "unauthorized"}`))
			return
		}

		// usr, err := user.Get(r.Context(), claims.Subject)
		// if err != nil {
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	w.Write([]byte(`{"error": "failed to get user"}`))
		// 	return
		// }
		next.ServeHTTP(w, r)
	})
}

// func RegisterRoutes(e *echo.Echo) {
// 	e.GET("auth-callback", auth.Callback)
// 	e.GET("/login", auth.ClerkSignIn)
// 	e.GET("/logout", echo.WrapHandler(protectedWithAuth(http.HandlerFunc(echoToHTTPHandler(auth.SignOut)))))
// 	e.GET("/create-session", echo.WrapHandler(protectedWithAuth(http.HandlerFunc(echoToHTTPHandler(sessions.CreateSession)))))
// 	e.GET("/end-session", echo.WrapHandler(protectedWithAuth(http.HandlerFunc(echoToHTTPHandler(sessions.EndSession)))))
// 	e.GET("/join-session", echo.WrapHandler(protectedWithAuth(http.HandlerFunc(echoToHTTPHandler(sessions.JoinSession)))))
// 	e.GET("/get-sessions", echo.WrapHandler(protectedWithAuth(http.HandlerFunc(echoToHTTPHandler(sessions.GetSessions)))))
// 	e.GET("/play-song", echo.WrapHandler(protectedWithAuth(http.HandlerFunc(echoToHTTPHandler(sessions.PlaySong)))))
// }

func RegisterRoutes(e *echo.Echo) {
	e.GET("auth-callback", auth.Callback)
	e.GET("/login", auth.SpotifySignIn)
	e.GET("/logout", auth.SignOut)
	e.GET("/create-session", sessions.CreateSession)
	e.GET("/end-session", sessions.EndSession)
	e.GET("/join-session", sessions.JoinSession)
	e.GET("/get-sessions", sessions.GetSessions)
	e.GET("/play-song", sessions.PlaySong)
}

func SpotifyUserMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if auth.AdminClient != nil {
			c.Set("user", auth.AdminClient)
		}
		return next(c)
	}
}

func GetUserFromContext(c echo.Context) (*spotify.PrivateUser, error) {
	// TODO : fix error when user is nil (occurs when user is not authenticated)
	userClient, ok := c.Get("user").(*spotify.PrivateUser)
	if !ok {
		return nil, fmt.Errorf("user not found in context")
	}
	fmt.Println("User found in context:", userClient.Email)
	return userClient, nil
}

func echoToHTTPHandler(echoHandler echo.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := echo.New().NewContext(r, w)
		if err := echoHandler(c); err != nil {
			c.Error(err)
		}
	}
}
