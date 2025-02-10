package service

import (
	"context"

	"github.com/arganaphang/openmusic/internal/entity"
	"github.com/arganaphang/openmusic/internal/repository"
)

type SongService interface {
	GetAll(ctx context.Context) ([]entity.Song, error)
	GetByID(ctx context.Context, id int) (*entity.Song, error)
	Create(ctx context.Context, song entity.Song) (*entity.Song, error)
	Update(ctx context.Context, id int, song entity.Song) (*entity.Song, error)
	Delete(ctx context.Context, id int) error
}

type songService struct {
	Repositories repository.Repositories
}

func NewSongService(repositories repository.Repositories) SongService {
	return &songService{
		Repositories: repositories,
	}
}

func (s songService) GetAll(ctx context.Context) ([]entity.Song, error) {
	return s.Repositories.SongRepository.GetAll(ctx)
}

func (s songService) GetByID(ctx context.Context, id int) (*entity.Song, error) {
	return s.Repositories.SongRepository.GetByID(ctx, id)
}

func (s songService) Create(ctx context.Context, album entity.Song) (*entity.Song, error) {
	return s.Repositories.SongRepository.Create(ctx, album)
}

func (s songService) Update(ctx context.Context, id int, album entity.Song) (*entity.Song, error) {
	return s.Repositories.SongRepository.Update(ctx, id, album)
}

func (s songService) Delete(ctx context.Context, id int) error {
	return s.Repositories.SongRepository.Delete(ctx, id)
}
