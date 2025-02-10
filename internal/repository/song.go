package repository

import (
	"context"

	"github.com/arganaphang/openmusic/internal/entity"
)

type SongRepository interface {
	GetAll(ctx context.Context) ([]entity.Song, error)
	GetByID(ctx context.Context, id int) (*entity.Song, error)
	Create(ctx context.Context, song entity.Song) (*entity.Song, error)
	Update(ctx context.Context, id int, song entity.Song) (*entity.Song, error)
	Delete(ctx context.Context, id int) error
}

type songRepository struct{}

func NewSongRepository() SongRepository {
	return &songRepository{}
}

func (r songRepository) GetAll(ctx context.Context) ([]entity.Song, error) {
	panic("not implemented") // TODO: Implement
}

func (r songRepository) GetByID(ctx context.Context, id int) (*entity.Song, error) {
	panic("not implemented") // TODO: Implement
}

func (r songRepository) Create(ctx context.Context, song entity.Song) (*entity.Song, error) {
	panic("not implemented") // TODO: Implement
}

func (r songRepository) Update(ctx context.Context, id int, song entity.Song) (*entity.Song, error) {
	panic("not implemented") // TODO: Implement
}

func (r songRepository) Delete(ctx context.Context, id int) error {
	panic("not implemented") // TODO: Implement
}
