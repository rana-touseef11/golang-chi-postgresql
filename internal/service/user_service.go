package service

import (
	"context"
	"time"

	"github.com/rana-touseef11/go-chi-postgresql/internal/dto"
	"github.com/rana-touseef11/go-chi-postgresql/internal/middleware"
	"github.com/rana-touseef11/go-chi-postgresql/internal/model"
	"github.com/rana-touseef11/go-chi-postgresql/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(r *repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) GetAll(ctx context.Context) ([]model.User, error) {
	return s.repo.GetAll(ctx)
}

func (s *UserService) Login(ctx context.Context, req dto.UserLoginRequest) (*model.User, string, error) {
	u := model.User{
		Email: req.Eamil,
	}

	user, err := s.repo.Login(ctx, u)
	if err != nil {
		return nil, "", err
	}

	token, err := middleware.JWTSign(user.ID.String(), time.Hour, "user")
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (s *UserService) Create(ctx context.Context, req dto.CreateUserRequest) (*model.User, error) {
	u := model.User{
		Name:  req.Name,
		Email: req.Email,
		Phone: req.Phone,
	}
	return s.repo.Create(ctx, u)
}

func (s *UserService) GetById(ctx context.Context, id string) (*model.User, error) {
	return s.repo.GetById(ctx, id)
}

func (s *UserService) Update(ctx context.Context, id string, req dto.UpdateUserRequest) (*model.User, error) {
	user, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	user.UpdateUserFromDTO(req)

	return s.repo.Update(ctx, id, *user)
}

func (s *UserService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
