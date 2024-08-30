package main

import (
	"log"

	todoapp "github.com/1tututu1"
	"github.com/1tututu1/pkg/handler"
	"github.com/1tututu1/pkg/repository"
	"github.com/1tututu1/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todoapp.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running")
	}
}
