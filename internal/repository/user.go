package repository

import (
	"context"
	"fmt"

	"github.com/arganaphang/openmusic/internal/dto"
	"github.com/arganaphang/openmusic/internal/entity"
	"github.com/arganaphang/openmusic/pkg"
	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type UserRepository interface {
	GetAll(ctx context.Context) ([]entity.User, error)
	GetByID(ctx context.Context, id string) (*entity.User, error)
	GetByUsername(ctx context.Context, username string) (*entity.User, error)
	Create(ctx context.Context, data dto.UserCreateRequest) (*entity.User, error)
	Update(ctx context.Context, id string, data dto.UserUpdateRequest) (*entity.User, error)
	Delete(ctx context.Context, id string) error
}

type userRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(DB *sqlx.DB) UserRepository {
	return &userRepository{
		DB: DB,
	}
}

func (r userRepository) GetAll(ctx context.Context) ([]entity.User, error) {
	sql, _, _ := goqu.From(entity.TABLE_USERS).ToSQL()
	users := []entity.User{}
	if err := r.DB.Select(&users, sql); err != nil {
		return nil, err
	}
	return users, nil
}

func (r userRepository) GetByID(ctx context.Context, id string) (*entity.User, error) {
	sql, _, _ := goqu.From(entity.TABLE_USERS).Where(goqu.C("id").Eq(id)).ToSQL()
	user := entity.User{}
	if err := r.DB.Get(&user, sql); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r userRepository) GetByUsername(ctx context.Context, username string) (*entity.User, error) {
	sql, _, _ := goqu.From(entity.TABLE_USERS).Where(goqu.C("username").Eq(username)).ToSQL()
	user := entity.User{}
	if err := r.DB.Get(&user, sql); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r userRepository) Create(ctx context.Context, data dto.UserCreateRequest) (*entity.User, error) {
	password, err := pkg.HashCreate(data.Password)
	if err != nil {
		return nil, err
	}
	id := fmt.Sprintf("user-%s", gonanoid.Must())
	sql, _, _ := goqu.Insert(entity.TABLE_USERS).
		Cols("id", "fullname", "username", "password").
		Vals(goqu.Vals{id, data.Fullname, data.Username, password}).
		Returning("*").
		ToSQL()
	rows, err := r.DB.Queryx(sql)
	if err != nil {
		return nil, err
	}
	var user entity.User
	if rows.Next() {
		if err := rows.StructScan(&user); err != nil {
			return nil, err
		}
	}
	return &user, nil
}

func (r userRepository) Update(ctx context.Context, id string, data dto.UserUpdateRequest) (*entity.User, error) {
	password, err := pkg.HashCreate(data.Password)
	if err != nil {
		return nil, err
	}
	sql, _, err := goqu.Update(entity.TABLE_USERS).
		Set(goqu.Record{
			"fullname": data.Fullname,
			"username": data.Username,
			"password": password,
		}).
		Where(goqu.C("id").Eq(id)).
		Returning("*").
		ToSQL()
	if err != nil {
		return nil, err
	}
	rows, err := r.DB.Queryx(sql)
	if err != nil {
		return nil, err
	}
	var user entity.User
	if rows.Next() {
		if err := rows.StructScan(&user); err != nil {
			return nil, err
		}
	}
	return &user, nil
}

func (r userRepository) Delete(ctx context.Context, id string) error {
	sql, _, _ := goqu.Delete(entity.TABLE_USERS).
		Where(goqu.C("id").Eq(id)).
		ToSQL()
	if _, err := r.DB.Exec(sql); err != nil {
		return err
	}
	return nil
}
