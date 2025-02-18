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
	Passcode   string
	Lock       sync.Mutex
}

type QueueSongRequest struct {
	SessionID int    `json:"session_id"`
	SongID    string `json:"song_id"`
}

type JoinSessionRequest struct {
	SessionID int `json:"session_id"`
}
