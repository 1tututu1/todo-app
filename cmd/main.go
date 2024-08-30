package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	todoapp "github.com/1tututu1"
	"github.com/1tututu1/pkg/handler"
	"github.com/1tututu1/pkg/repository"
	"github.com/1tututu1/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todoapp.Server)
	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running")
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
