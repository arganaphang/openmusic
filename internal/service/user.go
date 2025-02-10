package service

import (
	"context"

	"github.com/arganaphang/openmusic/internal/entity"
	"github.com/arganaphang/openmusic/internal/repository"
)

type UserService interface {
	GetAll(ctx context.Context) ([]entity.User, error)
	GetByID(ctx context.Context, id int) (*entity.User, error)
	Create(ctx context.Context, user entity.User) (*entity.User, error)
	Update(ctx context.Context, id int, user entity.User) (*entity.User, error)
	Delete(ctx context.Context, id int) error
}

type userService struct {
	Repositories *repository.Repositories
}

func NewUserService(repositories *repository.Repositories) UserService {
	return &userService{
		Repositories: repositories,
	}
}

func (s userService) GetAll(ctx context.Context) ([]entity.User, error) {
	return s.Repositories.UserRepository.GetAll(ctx)
}

func (s userService) GetByID(ctx context.Context, id int) (*entity.User, error) {
	return s.Repositories.UserRepository.GetByID(ctx, id)
}

func (s userService) Create(ctx context.Context, album entity.User) (*entity.User, error) {
	return s.Repositories.UserRepository.Create(ctx, album)
}

func (s userService) Update(ctx context.Context, id int, album entity.User) (*entity.User, error) {
	return s.Repositories.UserRepository.Update(ctx, id, album)
}

func (s userService) Delete(ctx context.Context, id int) error {
	return s.Repositories.UserRepository.Delete(ctx, id)
}
