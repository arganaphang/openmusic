package service

import (
	"context"

	"github.com/arganaphang/openmusic/internal/dto"
	"github.com/arganaphang/openmusic/internal/entity"
	"github.com/arganaphang/openmusic/internal/repository"
	"github.com/arganaphang/openmusic/pkg"
)

type UserService interface {
	GetAll(ctx context.Context) ([]entity.User, error)
	GetByID(ctx context.Context, id string) (*entity.User, error)
	GetByUsername(ctx context.Context, username string) (*entity.User, error)
	Create(ctx context.Context, data dto.UserCreateRequest) (*entity.User, error)
	Update(ctx context.Context, id string, data dto.UserUpdateRequest) (*entity.User, error)
	Delete(ctx context.Context, id string) error

	Register(ctx context.Context, data dto.RegisterRequest) (*entity.UserJWT, error)
	Login(ctx context.Context, data dto.LoginRequest) (*entity.UserJWT, error)
	RefreshToken(ctx context.Context, data dto.RefreshTokenRequest) (*entity.UserJWT, error)
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

func (s userService) GetByID(ctx context.Context, id string) (*entity.User, error) {
	return s.Repositories.UserRepository.GetByID(ctx, id)
}

func (s userService) GetByUsername(ctx context.Context, username string) (*entity.User, error) {
	return s.Repositories.UserRepository.GetByUsername(ctx, username)
}

func (s userService) Create(ctx context.Context, data dto.UserCreateRequest) (*entity.User, error) {
	return s.Repositories.UserRepository.Create(ctx, data)
}

func (s userService) Update(ctx context.Context, id string, data dto.UserUpdateRequest) (*entity.User, error) {
	return s.Repositories.UserRepository.Update(ctx, id, data)
}

func (s userService) Delete(ctx context.Context, id string) error {
	return s.Repositories.UserRepository.Delete(ctx, id)
}

func (s userService) Register(ctx context.Context, data dto.RegisterRequest) (*entity.UserJWT, error) {
	_, err := s.Repositories.UserRepository.Create(ctx, dto.UserCreateRequest{
		Fullname: data.Fullname,
		Username: data.Username,
		Password: data.Password,
	})
	if err != nil {
		return nil, err
	}
	// TODO: Create TOKEN
	return nil, nil
}

func (s userService) Login(ctx context.Context, data dto.LoginRequest) (*entity.UserJWT, error) {
	user, err := s.Repositories.UserRepository.GetByUsername(ctx, data.Username)
	if err != nil {
		return nil, err
	}
	if !pkg.HashCompare(data.Password, user.Password) {
		return nil, pkg.ErrUnauthorized
	}
	token, err := pkg.JWTEncode(*user)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (s userService) RefreshToken(ctx context.Context, data dto.RefreshTokenRequest) (*entity.UserJWT, error) {
	token, err := pkg.JWTRefresh(data.AccessToken, data.RefreshToken)
	if err != nil {
		return nil, err
	}
	return token, nil
}
