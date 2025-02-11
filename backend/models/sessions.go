package models

import (
	"sync"

	"github.com/zmb3/spotify"
)

type Session struct {
	ID         int
	Admin      *SpotifyUser
	UserIDs    []string
	SongQueue  []spotify.URI
	AdminToken string
	Lock       sync.Mutex
}
