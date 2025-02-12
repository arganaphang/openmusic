package service

import (
	"context"

	"github.com/arganaphang/openmusic/internal/dto"
	"github.com/arganaphang/openmusic/internal/entity"
	"github.com/arganaphang/openmusic/internal/repository"
)

type SongService interface {
	GetAll(ctx context.Context, params dto.SongGetAllRequest) ([]entity.Song, error)
	GetByID(ctx context.Context, id string) (*entity.Song, error)
	GetByAlbumID(ctx context.Context, albumID string) ([]entity.Song, error)
	Create(ctx context.Context, data dto.SongCreateRequest) (*entity.Song, error)
	Update(ctx context.Context, id string, data dto.SongUpdateRequest) (*entity.Song, error)
	Delete(ctx context.Context, id string) error
}

type songService struct {
	Repositories *repository.Repositories
}

func NewSongService(repositories *repository.Repositories) SongService {
	return &songService{
		Repositories: repositories,
	}
}

func (s songService) GetAll(ctx context.Context, params dto.SongGetAllRequest) ([]entity.Song, error) {
	return s.Repositories.SongRepository.GetAll(ctx, params)
}

func (s songService) GetByID(ctx context.Context, id string) (*entity.Song, error) {
	return s.Repositories.SongRepository.GetByID(ctx, id)
}

func (s songService) GetByAlbumID(ctx context.Context, albumID string) ([]entity.Song, error) {
	return s.Repositories.SongRepository.GetByAlbumID(ctx, albumID)
}

func (s songService) Create(ctx context.Context, data dto.SongCreateRequest) (*entity.Song, error) {
	return s.Repositories.SongRepository.Create(ctx, data)
}

func (s songService) Update(ctx context.Context, id string, data dto.SongUpdateRequest) (*entity.Song, error) {
	return s.Repositories.SongRepository.Update(ctx, id, data)
}

func (s songService) Delete(ctx context.Context, id string) error {
	return s.Repositories.SongRepository.Delete(ctx, id)
}
