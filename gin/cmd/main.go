package main

import (
	"gin/internal/handler"
	"gin/internal/repositories"
	"gin/internal/services"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	url := "postgresql://cinemakino_user:LqrK9KMsCLCwIYg7MfmGl4me1oK1BCim@dpg-csgfn0dumphs73b4t2ng-a.frankfurt-postgres.render.com/cinemakino"
	db, err := repositories.NewPostgresDB(url)
	if err != nil {
		logrus.Fatal("Error from db: ", err)
		return
	}

	repository := repositories.NewRepository(db)
	service := services.NewServices(repository)
	handler := handler.NewHandler(service)

	server := &http.Server{
		Addr:    ":5050",
		Handler: handler.InitRoutes(),
	}

	logrus.Info("Server starting on port 5050...")

	if err := server.ListenAndServe(); err != nil {
		logrus.Fatal(err)
	}

}
