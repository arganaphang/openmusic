package repository

import (
	"context"

	"github.com/arganaphang/openmusic/internal/entity"
	"github.com/arganaphang/openmusic/internal/repository/queries"
)

type PlaylistRepository interface {
	GetAll(ctx context.Context) ([]entity.Playlist, error)
	GetByID(ctx context.Context, id int) (*entity.Playlist, error)
	Create(ctx context.Context, playlist entity.Playlist) (*entity.Playlist, error)
	Update(ctx context.Context, id int, playlist entity.Playlist) (*entity.Playlist, error)
	Delete(ctx context.Context, id int) error
}

type playlistRepository struct {
	Queries *queries.Queries
}

func NewPlaylistRepository(queries *queries.Queries) PlaylistRepository {
	return &playlistRepository{
		Queries: queries,
	}
}

func (r playlistRepository) GetAll(ctx context.Context) ([]entity.Playlist, error) {
	panic("not implemented") // TODO: Implement
}

func (r playlistRepository) GetByID(ctx context.Context, id int) (*entity.Playlist, error) {
	panic("not implemented") // TODO: Implement
}

func (r playlistRepository) Create(ctx context.Context, playlist entity.Playlist) (*entity.Playlist, error) {
	panic("not implemented") // TODO: Implement
}

func (r playlistRepository) Update(ctx context.Context, id int, playlist entity.Playlist) (*entity.Playlist, error) {
	panic("not implemented") // TODO: Implement
}

func (r playlistRepository) Delete(ctx context.Context, id int) error {
	panic("not implemented") // TODO: Implement
}
