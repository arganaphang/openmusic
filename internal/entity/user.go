package entity

import "time"

const TABLE_USERS = "users"

type UserJWT struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type User struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u User) JWTCreateToken() (*UserJWT, error) {
	panic("unimplemented")
}

func (u User) JWTValidateToken() (*UserJWT, error) {
	panic("unimplemented")
}
