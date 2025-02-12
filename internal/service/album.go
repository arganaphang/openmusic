package service

import (
	"context"

	"github.com/arganaphang/openmusic/internal/dto"
	"github.com/arganaphang/openmusic/internal/entity"
	"github.com/arganaphang/openmusic/internal/repository"
)

type AlbumService interface {
	GetAll(ctx context.Context) ([]entity.Album, error)
	GetByID(ctx context.Context, id string) (*entity.Album, error)
	Create(ctx context.Context, data dto.AlbumCreateRequest) (*entity.Album, error)
	Update(ctx context.Context, id string, data dto.AlbumUpdateRequest) (*entity.Album, error)
	Delete(ctx context.Context, id string) error
}

type albumService struct {
	Repositories *repository.Repositories
}

func NewAlbumService(repositories *repository.Repositories) AlbumService {
	return &albumService{
		Repositories: repositories,
	}
}

func (s albumService) GetAll(ctx context.Context) ([]entity.Album, error) {
	return s.Repositories.AlbumRepository.GetAll(ctx)
}

func (s albumService) GetByID(ctx context.Context, id string) (*entity.Album, error) {
	return s.Repositories.AlbumRepository.GetByID(ctx, id)
}

func (s albumService) Create(ctx context.Context, data dto.AlbumCreateRequest) (*entity.Album, error) {
	return s.Repositories.AlbumRepository.Create(ctx, data)
}

func (s albumService) Update(ctx context.Context, id string, data dto.AlbumUpdateRequest) (*entity.Album, error) {
	return s.Repositories.AlbumRepository.Update(ctx, id, data)
}

func (s albumService) Delete(ctx context.Context, id string) error {
	return s.Repositories.AlbumRepository.Delete(ctx, id)
}
