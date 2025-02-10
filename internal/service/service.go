package service

import "github.com/arganaphang/openmusic/internal/repository"

type Services struct {
	AlbumService    AlbumService
	PlaylistService PlaylistService
	SongService     SongService
	UserService     UserService
}

func NewServices(repositories *repository.Repositories) *Services {
	return &Services{
		AlbumService:    NewAlbumService(repositories),
		PlaylistService: NewPlaylistService(repositories),
		SongService:     NewSongService(repositories),
		UserService:     NewUserService(repositories),
	}
}
