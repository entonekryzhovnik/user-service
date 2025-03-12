package main

import (
	"github.com/entonekryzhovnik/user-service/config"
	"github.com/entonekryzhovnik/user-service/internal/repository"
	"github.com/entonekryzhovnik/user-service/internal/server"
	"github.com/entonekryzhovnik/user-service/internal/service"
)

func main() {
	db := config.InitDB()
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	server.StartGRPCServer(userService)
}
