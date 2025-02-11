package dto

import "github.com/arganaphang/openmusic/internal/entity"

type SongGetAllResponse struct {
	Success bool          `json:"success"`
	Message string        `json:"message"`
	Data    []entity.Song `json:"data"`
}

type SongGetByIDResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Data    *entity.Song `json:"data"`
}

type SongCreateRequest struct {
	Title     string `json:"title" binding:"required"`
	Year      int16  `json:"year" binding:"required"`
	Genre     string `json:"genre" binding:"required"`
	Performer string `json:"performer" binding:"required"`
	Duration  int16  `json:"duration" binding:"required"`
	AlbumID   string `json:"album_id" binding:"required"`
}

type SongCreateResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Data    *entity.Song `json:"data"`
}

type SongUpdateRequest struct {
	Title     string `json:"title" binding:"required"`
	Year      int16  `json:"year" binding:"required"`
	Genre     string `json:"genre" binding:"required"`
	Performer string `json:"performer" binding:"required"`
	Duration  int16  `json:"duration" binding:"required"`
	AlbumID   string `json:"album_id" binding:"required"`
}

type SongUpdateResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Data    *entity.Song `json:"data"`
}

type SongDeleteResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
