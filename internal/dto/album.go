package dto

import "github.com/arganaphang/openmusic/internal/entity"

type AlbumGetAllResponse struct {
	Status  string                  `json:"status"`
	Message string                  `json:"message"`
	Data    AlbumGetAllResponseData `json:"data"`
}

type AlbumGetAllResponseData struct {
	Albums []entity.Album `json:"albums"`
}

type AlbumGetByIDResponse struct {
	Status  string                   `json:"status"`
	Message string                   `json:"message"`
	Data    AlbumGetByIDResponseData `json:"data"`
}

type AlbumGetByIDResponseData struct {
	Album *entity.AlbumWithSongs `json:"album"`
}

type AlbumCreateRequest struct {
	Name string `json:"name" binding:"required"`
	Year int16  `json:"year" binding:"required"`
}

type AlbumCreateResponse struct {
	Status  string                  `json:"status"`
	Message string                  `json:"message"`
	Data    AlbumCreateResponseData `json:"data"`
}

type AlbumCreateResponseData struct {
	Album *entity.Album `json:"album"`
}

type AlbumUpdateRequest struct {
	Name string `json:"name" binding:"required"`
	Year int16  `json:"year" binding:"required"`
}

type AlbumUpdateResponse struct {
	Status  string                  `json:"status"`
	Message string                  `json:"message"`
	Data    AlbumUpdateResponseData `json:"data"`
}

type AlbumUpdateResponseData struct {
	Album *entity.Album `json:"album"`
}

type AlbumDeleteResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
