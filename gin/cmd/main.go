package main

import (
	"gin/internal/config"
	"gin/internal/handler"
	"gin/internal/repositories"
	"gin/internal/services"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	cfg, err := config.InitConfig("../config.yml")
	if err != nil {
		panic(err)
	}

	db, err := repositories.NewPostgresDB(cfg.DB.URL)
	if err != nil {
		logrus.Fatal("Error from db: ", err)
		return
	}

	repository := repositories.NewRepository(db)
	service := services.NewServices(repository, cfg)
	handler := handler.NewHandler(service)

	server := &http.Server{
		Addr:    ":" + cfg.Server.Port,
		Handler: handler.InitRoutes(),
	}

	logrus.Info("Server starting on port 5050...")

	if err := server.ListenAndServe(); err != nil {
		logrus.Fatal(err)
	}
}
