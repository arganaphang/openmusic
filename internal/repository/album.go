package repository

import (
	"context"
	"fmt"

	"github.com/arganaphang/openmusic/internal/entity"
	"github.com/arganaphang/openmusic/internal/repository/queries"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type AlbumRepository interface {
	GetAll(ctx context.Context) ([]entity.Album, error)
	GetByID(ctx context.Context, id string) (*entity.Album, error)
	Create(ctx context.Context, album entity.Album) (*entity.Album, error)
	Update(ctx context.Context, id string, album entity.Album) (*entity.Album, error)
	Delete(ctx context.Context, id string) error
}

type albumRepository struct {
	Queries *queries.Queries
}

func NewAlbumRepository(queries *queries.Queries) AlbumRepository {
	return &albumRepository{
		Queries: queries,
	}
}

func (r albumRepository) GetAll(ctx context.Context) ([]entity.Album, error) {
	results, err := r.Queries.GetAlbums(ctx)
	if err != nil {
		return nil, err
	}
	albums := []entity.Album{}
	for _, row := range results {
		albums = append(albums, entity.Album{
			ID:   row.ID,
			Name: row.Name,
			Year: row.Year,
		})
	}
	return albums, nil
}

func (r albumRepository) GetByID(ctx context.Context, id string) (*entity.Album, error) {
	row, err := r.Queries.GetAlbumByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &entity.Album{
		ID:   row.ID,
		Name: row.Name,
		Year: row.Year,
	}, nil
}

func (r albumRepository) Create(ctx context.Context, album entity.Album) (*entity.Album, error) {
	result, err := r.Queries.CreateAlbum(ctx, queries.CreateAlbumParams{
		ID:   fmt.Sprintf("album-%s", gonanoid.Must()),
		Name: album.Name,
		Year: album.Year,
	})
	if err != nil {
		return nil, err
	}
	return &entity.Album{
		ID:   result.ID,
		Name: result.Name,
		Year: result.Year,
	}, nil
}

func (r albumRepository) Update(ctx context.Context, id string, album entity.Album) (*entity.Album, error) {
	result, err := r.Queries.UpdateAlbum(ctx, queries.UpdateAlbumParams{
		ID:   id,
		Name: album.Name,
		Year: album.Year,
	})
	if err != nil {
		return nil, err
	}
	return &entity.Album{
		ID:   result.ID,
		Name: result.Name,
		Year: result.Year,
	}, nil
}

func (r albumRepository) Delete(ctx context.Context, id string) error {
	return r.Queries.DeleteAlbum(ctx, id)
}
