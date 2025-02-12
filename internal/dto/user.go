package dto

import "github.com/arganaphang/openmusic/internal/entity"

type SongGetAllRequest struct {
	Title     string `json:"title" form:"title"`
	Performer string `json:"performer" form:"performer"`
}

type UserGetAllResponse struct {
	Status  string                 `json:"status"`
	Message string                 `json:"message"`
	Data    UserGetAllResponseData `json:"data"`
}

type UserGetAllResponseData struct {
	Users []entity.User `json:"users"`
}

type UserGetByIDResponse struct {
	Status  string                  `json:"status"`
	Message string                  `json:"message"`
	Data    UserGetByIDResponseData `json:"data"`
}

type UserGetByIDResponseData struct {
	User *entity.User `json:"users"`
}

type UserCreateRequest struct {
	Fullname string `json:"fullname" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserCreateResponse struct {
	Status  string                 `json:"status"`
	Message string                 `json:"message"`
	Data    UserCreateResponseData `json:"data"`
}

type UserCreateResponseData struct {
	User *entity.User `json:"user"`
}

type UserUpdateRequest struct {
	Fullname string `json:"fullname" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserUpdateResponse struct {
	Status  string                 `json:"status"`
	Message string                 `json:"message"`
	Data    UserUpdateResponseData `json:"data"`
}

type UserUpdateResponseData struct {
	User *entity.User `json:"users"`
}

type UserDeleteResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type RegisterRequest struct {
	Fullname string `json:"fullname" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterResponse struct {
	Status  string               `json:"status"`
	Message string               `json:"message"`
	Data    RegisterResponseData `json:"data"`
}

type RegisterResponseData struct {
	Token *entity.UserJWT `json:"token"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Status  string            `json:"status"`
	Message string            `json:"message"`
	Data    LoginResponseData `json:"data"`
}

type LoginResponseData struct {
	Token *entity.UserJWT `json:"token"`
}
