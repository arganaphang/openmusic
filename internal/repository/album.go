package repository

import (
	"context"
	"fmt"

	"github.com/arganaphang/openmusic/internal/entity"
	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
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
	DB *sqlx.DB
}

func NewAlbumRepository(DB *sqlx.DB) AlbumRepository {
	return &albumRepository{
		DB: DB,
	}
}

func (r albumRepository) GetAll(ctx context.Context) ([]entity.Album, error) {
	sql, _, _ := goqu.From(entity.TABLE_ALBUMS).ToSQL()
	albums := []entity.Album{}
	if err := r.DB.Select(&albums, sql); err != nil {
		return nil, err
	}
	return albums, nil
}

func (r albumRepository) GetByID(ctx context.Context, id string) (*entity.Album, error) {
	sql, _, _ := goqu.From(entity.TABLE_ALBUMS).Where(goqu.C("id").Eq(id)).ToSQL()
	album := entity.Album{}
	if err := r.DB.Get(&album, sql); err != nil {
		return nil, err
	}
	return &album, nil
}

func (r albumRepository) Create(ctx context.Context, album entity.Album) (*entity.Album, error) {
	album.ID = fmt.Sprintf("album-%s", gonanoid.Must())
	sql, _, _ := goqu.Insert(entity.TABLE_ALBUMS).
		Cols("id", "name", "year").
		Vals(goqu.Vals{album.ID, album.Name, album.Year}).
		ToSQL()
	if _, err := r.DB.Exec(sql); err != nil {
		return nil, err
	}
	return &album, nil
}

func (r albumRepository) Update(ctx context.Context, id string, album entity.Album) (*entity.Album, error) {
	album.ID = id
	sql, _, err := goqu.Update(entity.TABLE_ALBUMS).
		Set(goqu.Record{
			"name": album.Name,
			"year": album.Year,
		}).
		Where(goqu.C("id").Eq(album.ID)).
		ToSQL()
	if err != nil {
		return nil, err
	}
	if _, err := r.DB.Exec(sql); err != nil {
		return nil, err
	}
	return &album, nil
}

func (r albumRepository) Delete(ctx context.Context, id string) error {
	sql, _, _ := goqu.Delete(entity.TABLE_ALBUMS).
		Where(goqu.C("id").Eq(id)).
		ToSQL()
	if _, err := r.DB.Exec(sql); err != nil {
		return err
	}
	return nil
}
