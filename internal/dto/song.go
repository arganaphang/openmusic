package dto

import "github.com/arganaphang/openmusic/internal/entity"

type SongGetAllResponse struct {
	Status  string                 `json:"status"`
	Message string                 `json:"message"`
	Data    SongGetAllResponseData `json:"data"`
}

type SongGetAllResponseData struct {
	Songs []entity.Song `json:"songs"`
}

type SongGetByIDResponse struct {
	Status  string                  `json:"status"`
	Message string                  `json:"message"`
	Data    SongGetByIDResponseData `json:"data"`
}

type SongGetByIDResponseData struct {
	Song *entity.Song `json:"song"`
}

type SongCreateRequest struct {
	Title     string  `json:"title" binding:"required"`
	Year      int16   `json:"year" binding:"required"`
	Genre     string  `json:"genre" binding:"required"`
	Performer string  `json:"performer" binding:"required"`
	Duration  int16   `json:"duration" binding:"required"`
	AlbumID   *string `json:"album_id"`
}

type SongCreateResponse struct {
	Status  string                 `json:"status"`
	Message string                 `json:"message"`
	Data    SongCreateResponseData `json:"data"`
}

type SongCreateResponseData struct {
	Song *entity.Song `json:"song"`
}

type SongUpdateRequest struct {
	Title     string  `json:"title" binding:"required"`
	Year      int16   `json:"year" binding:"required"`
	Genre     string  `json:"genre" binding:"required"`
	Performer string  `json:"performer" binding:"required"`
	Duration  int16   `json:"duration" binding:"required"`
	AlbumID   *string `json:"album_id"`
}

type SongUpdateResponse struct {
	Status  string                 `json:"status"`
	Message string                 `json:"message"`
	Data    SongUpdateResponseData `json:"data"`
}

type SongUpdateResponseData struct {
	Song *entity.Song `json:"song"`
}

type SongDeleteResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
