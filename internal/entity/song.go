package entity

import "time"

type Song struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Year      int16     `json:"year"`
	Genre     string    `json:"genre"`
	Performer string    `json:"performer"`
	Duration  int16     `json:"duration"`
	AlbumID   string    `json:"album_id"`
	CreatedAt time.Time `json:"created_at"`
}
