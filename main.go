package main

import (
	"user-service/config"
	"user-service/internal/repository"
	"user-service/internal/server"
	"user-service/internal/service"
)

func main() {
	db := config.InitDB()
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	server.StartGRPCServer(userService)
}
