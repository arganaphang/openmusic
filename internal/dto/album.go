package dto

import "github.com/arganaphang/openmusic/internal/entity"

type AlbumGetAllResponse struct {
	Success bool           `json:"success"`
	Message string         `json:"message"`
	Data    []entity.Album `json:"data"`
}

type AlbumGetByIDResponse struct {
	Success bool          `json:"success"`
	Message string        `json:"message"`
	Data    *entity.Album `json:"data"`
}

type AlbumCreateRequest struct {
	Name string `json:"name" binding:"required"`
	Year int16  `json:"year" binding:"required"`
}

type AlbumCreateResponse struct {
	Success bool          `json:"success"`
	Message string        `json:"message"`
	Data    *entity.Album `json:"data"`
}

type AlbumUpdateRequest struct {
	Name string `json:"name" binding:"required"`
	Year int16  `json:"year" binding:"required"`
}

type AlbumUpdateResponse struct {
	Success bool          `json:"success"`
	Message string        `json:"message"`
	Data    *entity.Album `json:"data"`
}

type AlbumDeleteResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
