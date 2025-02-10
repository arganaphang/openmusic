package dto

type AlbumCreateRequest struct {
	Name string `json:"name" binding:"required"`
	Year int16  `json:"year" binding:"required"`
}

type AlbumUpdateRequest struct {
	Name string `json:"name" binding:"required"`
	Year int16  `json:"year" binding:"required"`
}
