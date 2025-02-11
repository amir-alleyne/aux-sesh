package models

import (
	"time"

	"github.com/zmb3/spotify"
)

type SpotifyUser struct {
	ID           string
	AccessToken  string
	RefreshToken string
	Expiry       time.Time
	Client       *spotify.Client
}

type User struct {
	ID        string `gorm:"primaryKey"` // Clerk User ID
	Email     string
	FirstName string
	LastName  string
	IsAdmin   bool
}
