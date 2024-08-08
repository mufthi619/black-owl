package main

import (
	"black-owl/internal/app"
	"black-owl/internal/config"
	"black-owl/internal/users/repository"
	"black-owl/internal/users/service"
)

func main() {

	appConfig := config.AppConfig{
		Database: config.DatabaseConfig{
			Host:     "localhost",
			Username: "startup",
			Password: "Startup123!",
			Database: "black_owl",
			Port:     "5432",
		},
	}
	db, err := app.NewDatabaseClient(appConfig.Database)
	if err != nil {
		panic("Failed to connect DB...")
	}

	allRepo := repository.NewInit(db, db)
	allService := service.NewInit(db, db)

}
