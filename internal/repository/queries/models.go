// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package queries

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Album struct {
	ID   string
	Name string
	Year int16
}

type Song struct {
	ID        string
	Title     string
	Year      int16
	Genre     string
	Performer string
	Duration  int16
	AlbumID   pgtype.Text
	CreatedAt pgtype.Timestamp
}

type User struct {
	ID        int32
	Name      string
	Email     string
	Password  string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}
