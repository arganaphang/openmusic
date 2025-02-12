package repository

import (
	"context"

	"github.com/arganaphang/openmusic/internal/entity"
	"github.com/jmoiron/sqlx"
)

type PlaylistRepository interface {
	GetAll(ctx context.Context) ([]entity.Playlist, error)
	GetByID(ctx context.Context, id int) (*entity.Playlist, error)
	Create(ctx context.Context, playlist entity.Playlist) (*entity.Playlist, error)
	Update(ctx context.Context, id int, playlist entity.Playlist) (*entity.Playlist, error)
	Delete(ctx context.Context, id int) error
}

type playlistRepository struct {
	DB *sqlx.DB
}

func NewPlaylistRepository(DB *sqlx.DB) PlaylistRepository {
	return &playlistRepository{
		DB: DB,
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
