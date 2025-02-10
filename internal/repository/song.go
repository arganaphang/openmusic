package repository

import (
	"context"

	"github.com/arganaphang/openmusic/internal/entity"
	"github.com/arganaphang/openmusic/internal/repository/queries"
)

type SongRepository interface {
	GetAll(ctx context.Context) ([]entity.Song, error)
	GetByID(ctx context.Context, id string) (*entity.Song, error)
	Create(ctx context.Context, song entity.Song) (*entity.Song, error)
	Update(ctx context.Context, id string, song entity.Song) (*entity.Song, error)
	Delete(ctx context.Context, id string) error
}

type songRepository struct {
	Queries *queries.Queries
}

func NewSongRepository(queries *queries.Queries) SongRepository {
	return &songRepository{
		Queries: queries,
	}
}

func (r songRepository) GetAll(ctx context.Context) ([]entity.Song, error) {
	panic("not implemented") // TODO: Implement
}

func (r songRepository) GetByID(ctx context.Context, id string) (*entity.Song, error) {
	panic("not implemented") // TODO: Implement
}

func (r songRepository) Create(ctx context.Context, song entity.Song) (*entity.Song, error) {
	panic("not implemented") // TODO: Implement
}

func (r songRepository) Update(ctx context.Context, id string, song entity.Song) (*entity.Song, error) {
	panic("not implemented") // TODO: Implement
}

func (r songRepository) Delete(ctx context.Context, id string) error {
	panic("not implemented") // TODO: Implement
}
