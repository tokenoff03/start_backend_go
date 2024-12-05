package services

import (
	"gin/internal/models"
	"gin/internal/repositories"
)

type Room interface {
	GetRoomById(id int) (models.Room, error)
	CreateRoom(room models.Room) (int, string)
}

type Service struct {
	Room
}

func NewServices(str *repositories.Storage) *Service {
	return &Service{
		Room: NewRoomService(str),
	}
}
