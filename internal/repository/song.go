package repository

import (
	"context"
	"fmt"

	"github.com/arganaphang/openmusic/internal/dto"
	"github.com/arganaphang/openmusic/internal/entity"
	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type SongRepository interface {
	GetAll(ctx context.Context, params dto.SongGetAllRequest) ([]entity.Song, error)
	GetByID(ctx context.Context, id string) (*entity.Song, error)
	GetByAlbumID(ctx context.Context, albumID string) ([]entity.Song, error)
	Create(ctx context.Context, data dto.SongCreateRequest) (*entity.Song, error)
	Update(ctx context.Context, id string, data dto.SongUpdateRequest) (*entity.Song, error)
	Delete(ctx context.Context, id string) error
}

type songRepository struct {
	DB *sqlx.DB
}

func NewSongRepository(DB *sqlx.DB) SongRepository {
	return &songRepository{
		DB: DB,
	}
}

func (r songRepository) GetAll(ctx context.Context, params dto.SongGetAllRequest) ([]entity.Song, error) {
	stmt := goqu.From(entity.TABLE_SONGS)
	if params.Title != "" {
		stmt = stmt.Where(goqu.C("title").ILike(fmt.Sprintf("%%%s%%", params.Title)))
	}
	if params.Performer != "" {
		stmt = stmt.Where(goqu.C("performer").ILike(fmt.Sprintf("%%%s%%", params.Performer)))
	}
	sql, _, _ := stmt.ToSQL()
	songs := []entity.Song{}
	if err := r.DB.Select(&songs, sql); err != nil {
		return nil, err
	}
	return songs, nil
}

func (r songRepository) GetByID(ctx context.Context, id string) (*entity.Song, error) {
	sql, _, _ := goqu.From(entity.TABLE_SONGS).Where(goqu.Ex{"id": id}).ToSQL()
	song := entity.Song{}
	if err := r.DB.Get(&song, sql); err != nil {
		return nil, err
	}
	return &song, nil
}

func (r songRepository) GetByAlbumID(ctx context.Context, albumID string) ([]entity.Song, error) {
	sql, _, _ := goqu.From(entity.TABLE_SONGS).Where(goqu.C("album_id").Eq(albumID)).ToSQL()
	songs := []entity.Song{}
	if err := r.DB.Select(&songs, sql); err != nil {
		return nil, err
	}
	return songs, nil
}

func (r songRepository) Create(ctx context.Context, data dto.SongCreateRequest) (*entity.Song, error) {
	song := entity.Song{
		ID:        fmt.Sprintf("song-%s", gonanoid.Must()),
		Title:     data.Title,
		Year:      data.Year,
		Genre:     data.Genre,
		Performer: data.Performer,
		Duration:  data.Duration,
		AlbumID:   data.AlbumID,
	}
	sql, _, _ := goqu.Insert(entity.TABLE_SONGS).
		Cols("id", "title", "year", "genre", "performer", "duration", "album_id").
		Vals(goqu.Vals{song.ID, song.Title, song.Year, song.Genre, song.Performer, song.Duration, song.AlbumID}).
		ToSQL()
	if _, err := r.DB.Exec(sql); err != nil {
		return nil, err
	}
	return &song, nil
}

func (r songRepository) Update(ctx context.Context, id string, data dto.SongUpdateRequest) (*entity.Song, error) {
	song := entity.Song{
		ID:        id,
		Title:     data.Title,
		Year:      data.Year,
		Genre:     data.Genre,
		Performer: data.Performer,
		Duration:  data.Duration,
		AlbumID:   data.AlbumID,
	}
	sql, _, _ := goqu.Update(entity.TABLE_SONGS).
		Set(goqu.Record{
			"title":     song.Title,
			"year":      song.Year,
			"genre":     song.Genre,
			"performer": song.Performer,
			"duration":  song.Duration,
			"album_id":  song.AlbumID,
		}).
		Where(goqu.C("id").Eq(song.ID)).
		ToSQL()
	if _, err := r.DB.Exec(sql); err != nil {
		return nil, err
	}
	return &song, nil
}

func (r songRepository) Delete(ctx context.Context, id string) error {
	sql, _, _ := goqu.Delete(entity.TABLE_SONGS).
		Where(goqu.C("id").Eq(id)).
		ToSQL()
	if _, err := r.DB.Exec(sql); err != nil {
		return err
	}
	return nil
}
