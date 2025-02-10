package service

import (
	"context"

	"github.com/arganaphang/openmusic/internal/entity"
	"github.com/arganaphang/openmusic/internal/repository"
)

type AlbumService interface {
	GetAll(ctx context.Context) ([]entity.Album, error)
	GetByID(ctx context.Context, id int) (*entity.Album, error)
	Create(ctx context.Context, album entity.Album) (*entity.Album, error)
	Update(ctx context.Context, id int, album entity.Album) (*entity.Album, error)
	Delete(ctx context.Context, id int) error
}

type albumService struct {
	Repositories repository.Repositories
}

func NewAlbumService(repositories repository.Repositories) AlbumService {
	return &albumService{
		Repositories: repositories,
	}
}

func (s albumService) GetAll(ctx context.Context) ([]entity.Album, error) {
	return s.Repositories.AlbumRepository.GetAll(ctx)
}

func (s albumService) GetByID(ctx context.Context, id int) (*entity.Album, error) {
	return s.Repositories.AlbumRepository.GetByID(ctx, id)
}

func (s albumService) Create(ctx context.Context, album entity.Album) (*entity.Album, error) {
	return s.Repositories.AlbumRepository.Create(ctx, album)
}

func (s albumService) Update(ctx context.Context, id int, album entity.Album) (*entity.Album, error) {
	return s.Repositories.AlbumRepository.Update(ctx, id, album)
}

func (s albumService) Delete(ctx context.Context, id int) error {
	return s.Repositories.AlbumRepository.Delete(ctx, id)
}
