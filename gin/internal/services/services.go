package services

import (
	"gin/internal/config"
	"gin/internal/models"
	"gin/internal/repositories"
)

type Room interface {
	GetRoomById(id int) (models.Room, error)
	CreateRoom(room models.RoomCreate) (int, error)
}

type Service struct {
	Room
}

func NewServices(repository *repositories.Repository, cfg *config.Config) *Service {
	return &Service{
		Room: NewRoomService(repository,cfg),
	}
}
