package entity

import "time"

const TABLE_USERS = "users"

type UserJWT struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type User struct {
	ID        string    `json:"id" db:"id"`
	Fullname  string    `json:"fullname" db:"fullname"`
	Username  string    `json:"username" db:"username"`
	Password  string    `json:"password" db:"password"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func (u User) JWTCreateToken() (*UserJWT, error) {
	panic("unimplemented")
}

func (u User) JWTValidateToken() (*UserJWT, error) {
	panic("unimplemented")
}
