package main

import (
	"mux/internal/handler"
	"mux/internal/repository"
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {
	mux := http.NewServeMux()

	storage := repository.NewStorage()
	handler := handler.NewHandler(storage)

	mux.HandleFunc("/room/", handler.RoomHandler)
	mux.HandleFunc("/room", handler.RoomHandler)

	server := &http.Server{
		Addr:    ":5050",
		Handler: mux,
	}

	logrus.Info("Server starting on port 5050...")

	if err := server.ListenAndServe(); err != nil {
		logrus.Fatal(err)
	}
}
