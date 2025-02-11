package dto

type SongCreateRequest struct {
	Title     string `json:"title" binding:"required"`
	Year      int16  `json:"year" binding:"required"`
	Genre     string `json:"genre" binding:"required"`
	Performer string `json:"performer" binding:"required"`
	Duration  int16  `json:"duration" binding:"required"`
	AlbumID   string `json:"album_id" binding:"required"`
}

type SongUpdateRequest struct {
	Title     string `json:"title" binding:"required"`
	Year      int16  `json:"year" binding:"required"`
	Genre     string `json:"genre" binding:"required"`
	Performer string `json:"performer" binding:"required"`
	Duration  int16  `json:"duration" binding:"required"`
	AlbumID   string `json:"album_id" binding:"required"`
}
