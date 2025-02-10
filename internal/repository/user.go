package repository

import (
	"context"

	"github.com/arganaphang/openmusic/internal/entity"
)

type UserRepository interface {
	GetAll(ctx context.Context) ([]entity.User, error)
	GetByID(ctx context.Context, id int) (*entity.User, error)
	Create(ctx context.Context, user entity.User) (*entity.User, error)
	Update(ctx context.Context, id int, user entity.User) (*entity.User, error)
	Delete(ctx context.Context, id int) error
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r userRepository) GetAll(ctx context.Context) ([]entity.User, error) {
	panic("not implemented") // TODO: Implement
}

func (r userRepository) GetByID(ctx context.Context, id int) (*entity.User, error) {
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
