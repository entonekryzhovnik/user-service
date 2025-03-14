package main

import (
	"github.com/entonekryzhovnik/user-service/config"
	"github.com/entonekryzhovnik/user-service/internal/controller"
	"github.com/entonekryzhovnik/user-service/internal/repository"
	"github.com/entonekryzhovnik/user-service/internal/server"
	"github.com/entonekryzhovnik/user-service/internal/service"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// Инициализация логера Zerolog
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Logger = zerolog.New(os.Stdout).With().
		Timestamp().
		Caller().
		Logger()

	log.Info().Msg("Starting user-service...")

	// Инициализация БД
	db := config.InitDB()
	defer db.Close()

	// Создание зависимостей
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	// Запуск gRPC-сервера
	server.StartGRPCServer(userController)
}
