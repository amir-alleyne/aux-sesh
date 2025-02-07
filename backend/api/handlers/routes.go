package handlers

import (
	"github.com/amir-alleyne/aux-sesh/backend/api/auth"
	"github.com/amir-alleyne/aux-sesh/backend/api/sessions"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("auth-callback", auth.Callback)
	e.GET("/login", auth.SignIn)
	e.GET("/logout", auth.SignOut)
	e.GET("/create-session", sessions.CreateSession)
	e.GET("/end-session", sessions.EndSession)
	e.GET("/join-session", sessions.JoinSession)
	e.GET("/get-sessions", sessions.GetSessions)
	e.GET("/play-song", sessions.PlaySong)
}
