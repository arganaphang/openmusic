package repository

import (
	"context"

	"github.com/arganaphang/openmusic/internal/entity"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	GetAll(ctx context.Context) ([]entity.User, error)
	GetByID(ctx context.Context, id int) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	Create(ctx context.Context, user entity.User) (*entity.User, error)
	Update(ctx context.Context, id int, user entity.User) (*entity.User, error)
	Delete(ctx context.Context, id int) error
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
	panic("not implemented") // TODO: Implement
}

func (r userRepository) GetByID(ctx context.Context, id int) (*entity.User, error) {
	panic("not implemented") // TODO: Implement
}

func (r userRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	panic("not implemented") // TODO: Implement
}

func (r userRepository) Create(ctx context.Context, user entity.User) (*entity.User, error) {
	panic("not implemented") // TODO: Implement
}

func (r userRepository) Update(ctx context.Context, id int, user entity.User) (*entity.User, error) {
	panic("not implemented") // TODO: Implement
}

func (r userRepository) Delete(ctx context.Context, id int) error {
	panic("not implemented") // TODO: Implement
}
