package service

import (
	"context"

	"github.com/arganaphang/openmusic/internal/entity"
	"github.com/arganaphang/openmusic/internal/repository"
)

type PlaylistService interface {
	GetAll(ctx context.Context) ([]entity.Playlist, error)
	GetByID(ctx context.Context, id int) (*entity.Playlist, error)
	Create(ctx context.Context, playlist entity.Playlist) (*entity.Playlist, error)
	Update(ctx context.Context, id int, playlist entity.Playlist) (*entity.Playlist, error)
	Delete(ctx context.Context, id int) error
}

type playlistService struct {
	Repositories repository.Repositories
}

func NewPlaylistService(repositories repository.Repositories) PlaylistService {
	return &playlistService{
		Repositories: repositories,
	}
}

func (s playlistService) GetAll(ctx context.Context) ([]entity.Playlist, error) {
	return s.Repositories.PlaylistRepository.GetAll(ctx)
}

func (s playlistService) GetByID(ctx context.Context, id int) (*entity.Playlist, error) {
	return s.Repositories.PlaylistRepository.GetByID(ctx, id)
}

func (s playlistService) Create(ctx context.Context, playlist entity.Playlist) (*entity.Playlist, error) {
	return s.Repositories.PlaylistRepository.Create(ctx, playlist)
}

func (s playlistService) Update(ctx context.Context, id int, playlist entity.Playlist) (*entity.Playlist, error) {
	return s.Repositories.PlaylistRepository.Update(ctx, id, playlist)
}

func (s playlistService) Delete(ctx context.Context, id int) error {
	return s.Repositories.PlaylistRepository.Delete(ctx, id)
}
