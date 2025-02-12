package entity

const TABLE_ALBUMS = "albums"

type Album struct {
	ID   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Year int16  `json:"year" db:"year"`
}

type AlbumWithSongs struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Year  int16  `json:"year"`
	Songs []Song `json:"songs"`
}
