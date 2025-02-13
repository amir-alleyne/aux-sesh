package main

import (
	"log"
	"os"

	"github.com/amir-alleyne/aux-sesh/backend/api/auth"
	"github.com/amir-alleyne/aux-sesh/backend/api/sessions"
	echoMiddleware "github.com/labstack/echo/v4/middleware"

	"github.com/joho/godotenv"

	"github.com/clerk/clerk-sdk-go/v2"

	echo "github.com/labstack/echo/v4"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file:", err)
	}
	clerkKey := os.Getenv("CLERK_SECRET_KEY")
	clerk.SetKey(clerkKey)

	if _, err := auth.SetAuth(); err != nil {
		log.Fatal("Error setting auth:", err)
	}

	e := echo.New()
	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	RegisterRoutes(e)

	if err := e.Start(":8080"); err != nil {
		log.Fatal("Error starting server:", err)
	}
}

func RegisterRoutes(e *echo.Echo) {
	e.GET("auth-callback", auth.Callback)
	e.GET("/login", auth.SpotifySignIn)
	e.GET("/logout", auth.SignOut)
	e.GET("/create-session", sessions.CreateSession)
	e.GET("/end-session", sessions.EndSession)
	e.GET("/join-session", sessions.JoinSession)
	e.GET("/get-sessions", sessions.GetSessions)
	e.POST("/queue-song", sessions.AddSongToQueue)
	e.GET("/play-song", sessions.PlaySong)
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

// func echoToHTTPHandler(echoHandler echo.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		c := echo.New().NewContext(r, w)
// 		if err := echoHandler(c); err != nil {
// 			c.Error(err)
// 		}
// 	}
// }

// func protectedWithAuth(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		_, ok := clerk.SessionClaimsFromContext(r.Context())
// 		if !ok {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			w.Write([]byte(`{"access": "unauthorized"}`))
// 			return
// 		}

// 		// usr, err := user.Get(r.Context(), claims.Subject)
// 		// if err != nil {
// 		// 	w.WriteHeader(http.StatusInternalServerError)
// 		// 	w.Write([]byte(`{"error": "failed to get user"}`))
// 		// 	return
// 		// }
// 		next.ServeHTTP(w, r)
// 	})
// }
