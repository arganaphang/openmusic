package repository

import (
	"context"

	"github.com/arganaphang/openmusic/internal/entity"
)

type AlbumRepository interface {
	GetAll(ctx context.Context) ([]entity.Album, error)
	GetByID(ctx context.Context, id int) (*entity.Album, error)
	Create(ctx context.Context, album entity.Album) (*entity.Album, error)
	Update(ctx context.Context, id int, album entity.Album) (*entity.Album, error)
	Delete(ctx context.Context, id int) error
}

type albumRepository struct{}

func NewAlbumRepository() AlbumRepository {
	return &albumRepository{}
}

func (r albumRepository) GetAll(ctx context.Context) ([]entity.Album, error) {
	panic("not implemented") // TODO: Implement
}

func (r albumRepository) GetByID(ctx context.Context, id int) (*entity.Album, error) {
	panic("not implemented") // TODO: Implement
}

func (r albumRepository) Create(ctx context.Context, album entity.Album) (*entity.Album, error) {
	panic("not implemented") // TODO: Implement
}

func (r albumRepository) Update(ctx context.Context, id int, album entity.Album) (*entity.Album, error) {
	panic("not implemented") // TODO: Implement
}

func (r albumRepository) Delete(ctx context.Context, id int) error {
	panic("not implemented") // TODO: Implement
}
