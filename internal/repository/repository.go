package repository

import "github.com/jmoiron/sqlx"

type Repositories struct {
	AlbumRepository    AlbumRepository
	PlaylistRepository PlaylistRepository
	SongRepository     SongRepository
	UserRepository     UserRepository
}

func NewRepositories(DB *sqlx.DB) *Repositories {
	return &Repositories{
		AlbumRepository:    NewAlbumRepository(DB),
		PlaylistRepository: NewPlaylistRepository(DB),
		SongRepository:     NewSongRepository(DB),
		UserRepository:     NewUserRepository(DB),
	}
}
