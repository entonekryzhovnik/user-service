package service

import (
	"context"
	"github.com/entonekryzhovnik/user-service/gen/go/userpb"
	"github.com/entonekryzhovnik/user-service/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
	userpb.UnimplementedUserServiceServer
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	user, err := s.repo.GetUser(req.Id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}
	return &userpb.GetUserResponse{User: user}, nil
}

func (s *UserService) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	id, err := s.repo.CreateUser(req.Email)
	if err != nil {
		return nil, err
	}
	return &userpb.CreateUserResponse{Id: id}, nil
}
