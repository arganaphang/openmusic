package entity

import "time"

const TABLE_SONGS = "songs"

type Song struct {
	ID        string    `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	Year      int16     `json:"year" db:"year"`
	Genre     string    `json:"genre" db:"genre"`
	Performer string    `json:"performer" db:"performer"`
	Duration  int16     `json:"duration" db:"duration"`
	AlbumID   *string   `json:"album_id" db:"album_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
