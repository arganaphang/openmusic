package repository

import (
	"github.com/arganaphang/openmusic/internal/repository/queries"
	"github.com/jackc/pgx/v5"
)

type Repositories struct {
	AlbumRepository    AlbumRepository
	PlaylistRepository PlaylistRepository
	SongRepository     SongRepository
	UserRepository     UserRepository
}

func NewRepositories(DB *pgx.Conn) *Repositories {
	queries := queries.New(DB)
	return &Repositories{
		AlbumRepository:    NewAlbumRepository(queries),
		PlaylistRepository: NewPlaylistRepository(queries),
		SongRepository:     NewSongRepository(queries),
		UserRepository:     NewUserRepository(queries),
	}
}
