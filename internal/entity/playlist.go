package entity

const TABLE_PLAYLISTS = "playlists"

type Playlist struct {
	ID          int
	Name        string
	Description string
}
