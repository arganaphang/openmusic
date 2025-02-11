package dto

import "github.com/arganaphang/openmusic/internal/entity"

type UserGetAllResponse struct {
	Success bool          `json:"success"`
	Message string        `json:"message"`
	Data    []entity.User `json:"data"`
}

type UserGetByIDResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Data    *entity.User `json:"data"`
}

type UserCreateRequest struct {
	Name     string `json:"name" binding:"required,email"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserCreateResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Data    *entity.User `json:"data"`
}

type UserUpdateRequest struct {
	Name     string `json:"name" binding:"required,email"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserUpdateResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Data    *entity.User `json:"data"`
}

type UserDeleteResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type RegisterRequest struct {
	Name     string `json:"name" binding:"required,email"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type RegisterResponse struct {
	Success bool            `json:"success"`
	Message string          `json:"message"`
	Data    *entity.UserJWT `json:"data"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Success bool            `json:"success"`
	Message string          `json:"message"`
	Data    *entity.UserJWT `json:"data"`
}
