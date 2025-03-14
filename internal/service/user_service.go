package service

import (
	"context"

	"github.com/entonekryzhovnik/user-service/internal/model"
	"github.com/entonekryzhovnik/user-service/internal/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, user model.User) (int64, error) // ✅ Принимаем model.User
	GetUser(ctx context.Context, id int64) (*model.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(ctx context.Context, user model.User) (int64, error) { // ✅ Принимаем model.User
	return s.repo.CreateUser(user) // ✅ Передаем структуру в репозиторий
}

func (s *userService) GetUser(ctx context.Context, id int64) (*model.User, error) {
	return s.repo.GetUser(id)
}
