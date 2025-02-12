package dto

import "github.com/arganaphang/openmusic/internal/entity"

type UserGetAllResponse struct {
	Status  string        `json:"success"`
	Message string        `json:"message"`
	Data    []entity.User `json:"data"`
}

type UserGetByIDResponse struct {
	Status  string       `json:"success"`
	Message string       `json:"message"`
	Data    *entity.User `json:"data"`
}

type UserCreateRequest struct {
	Name     string `json:"name" binding:"required,email"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserCreateResponse struct {
	Status  string       `json:"success"`
	Message string       `json:"message"`
	Data    *entity.User `json:"data"`
}

type UserUpdateRequest struct {
	Name     string `json:"name" binding:"required,email"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserUpdateResponse struct {
	Status  string       `json:"success"`
	Message string       `json:"message"`
	Data    *entity.User `json:"data"`
}

type UserDeleteResponse struct {
	Status  string `json:"success"`
	Message string `json:"message"`
}

type RegisterRequest struct {
	Name     string `json:"name" binding:"required,email"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type RegisterResponse struct {
	Status  string          `json:"success"`
	Message string          `json:"message"`
	Data    *entity.UserJWT `json:"data"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Status  string          `json:"success"`
	Message string          `json:"message"`
	Data    *entity.UserJWT `json:"data"`
}
