package repository

import (
	"context"
	"fmt"

	"github.com/arganaphang/openmusic/internal/entity"
	"github.com/arganaphang/openmusic/internal/repository/queries"
	"github.com/jackc/pgx/v5/pgtype"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type SongRepository interface {
	GetAll(ctx context.Context) ([]entity.Song, error)
	GetByID(ctx context.Context, id string) (*entity.Song, error)
	GetByAlbumID(ctx context.Context, albumID string) ([]entity.Song, error)
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
	results, err := r.Queries.GetSongs(ctx)
	if err != nil {
		return nil, err
	}
	songs := []entity.Song{}
	for _, result := range results {
		var albumID *string
		if result.AlbumID.Valid {
			albumID = &result.AlbumID.String
		}
		songs = append(songs, entity.Song{
			ID:        result.ID,
			Title:     result.Title,
			Performer: result.Performer,
			AlbumID:   albumID,
			Year:      result.Year,
			Genre:     result.Genre,
			Duration:  result.Duration,
			CreatedAt: result.CreatedAt.Time,
		})
	}
	return songs, nil
}

func (r songRepository) GetByID(ctx context.Context, id string) (*entity.Song, error) {
	result, err := r.Queries.GetSongByID(ctx, id)
	if err != nil {
		return nil, err
	}
	var albumID *string
	if result.AlbumID.Valid {
		albumID = &result.AlbumID.String
	}
	return &entity.Song{
		ID:        result.ID,
		Title:     result.Title,
		Performer: result.Performer,
		AlbumID:   albumID,
		Year:      result.Year,
		Genre:     result.Genre,
		Duration:  result.Duration,
		CreatedAt: result.CreatedAt.Time,
	}, nil
}

func (r songRepository) GetByAlbumID(ctx context.Context, albumID string) ([]entity.Song, error) {
	results, err := r.Queries.GetSongByAlbumID(ctx, pgtype.Text{
		Valid:  true,
		String: albumID,
	})
	if err != nil {
		return nil, err
	}
	songs := []entity.Song{}
	for _, result := range results {
		var albumID *string
		if result.AlbumID.Valid {
			albumID = &result.AlbumID.String
		}
		songs = append(songs, entity.Song{
			ID:        result.ID,
			Title:     result.Title,
			Performer: result.Performer,
			AlbumID:   albumID,
			Year:      result.Year,
			Genre:     result.Genre,
			Duration:  result.Duration,
			CreatedAt: result.CreatedAt.Time,
		})
	}
	return songs, nil
}

func (r songRepository) Create(ctx context.Context, song entity.Song) (*entity.Song, error) {
	var albumID string
	if song.AlbumID != nil {
		albumID = *song.AlbumID
	}
	result, err := r.Queries.CreateSong(ctx, queries.CreateSongParams{
		ID:        fmt.Sprintf("song-%s", gonanoid.Must()),
		Title:     song.Title,
		Year:      song.Year,
		Genre:     song.Genre,
		Performer: song.Performer,
		Duration:  song.Duration,
		AlbumID: pgtype.Text{
			Valid:  song.AlbumID != nil,
			String: albumID,
		},
	})
	if err != nil {
		return nil, err
	}
	return &entity.Song{
		ID:        result.ID,
		Title:     result.Title,
		Performer: result.Performer,
		AlbumID:   &albumID,
		Year:      result.Year,
		Genre:     result.Genre,
		Duration:  result.Duration,
		CreatedAt: result.CreatedAt.Time,
	}, nil
}

func (r songRepository) Update(ctx context.Context, id string, song entity.Song) (*entity.Song, error) {
	var albumID string
	if song.AlbumID != nil {
		albumID = *song.AlbumID
	}
	result, err := r.Queries.UpdateSong(ctx, queries.UpdateSongParams{
		ID:        id,
		Title:     song.Title,
		Year:      song.Year,
		Genre:     song.Genre,
		Performer: song.Performer,
		Duration:  song.Duration,
		AlbumID: pgtype.Text{
			Valid:  song.AlbumID != nil,
			String: albumID,
		},
	})
	if err != nil {
		return nil, err
	}
	return &entity.Song{
		ID:        result.ID,
		Title:     result.Title,
		Performer: result.Performer,
		AlbumID:   &albumID,
		Year:      result.Year,
		Genre:     result.Genre,
		Duration:  result.Duration,
		CreatedAt: result.CreatedAt.Time,
	}, nil
}

func (r songRepository) Delete(ctx context.Context, id string) error {
	return r.Queries.DeleteSong(ctx, id)
}
