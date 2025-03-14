package controller

import (
	"context"

	"github.com/entonekryzhovnik/user-service/gen/go/userpb"
	"github.com/entonekryzhovnik/user-service/internal/model"
	"github.com/entonekryzhovnik/user-service/internal/service"
)

type UserController struct {
	userpb.UnimplementedUserServiceServer
	service service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{service: service}
}

func (c *UserController) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	user := model.User{ // ✅ Создаем структуру model.User
		Email: req.Email,
	}
	id, err := c.service.CreateUser(ctx, user) // ✅ Передаем model.User в сервис
	if err != nil {
		return nil, err
	}
	return &userpb.CreateUserResponse{Id: id}, nil
}

func (c *UserController) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	user, err := c.service.GetUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}
	return &userpb.GetUserResponse{ // ✅ Конвертируем обратно в gRPC-структуру
		User: &userpb.User{
			Id:       user.ID,
			Email:    user.Email,
			CreateAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		},
	}, nil
}
