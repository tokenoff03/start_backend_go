package services

import (
	"gin/internal/models"
	"gin/internal/repositories"
)

type RoomService struct {
	repo *repositories.Repository
}

func NewRoomService(repo *repositories.Repository) *RoomService {
	return &RoomService{repo: repo}
}

func (s *RoomService) GetRoomById(id int) (models.Room, error) {
	return s.repo.GetRoomById(id)
}

func (s *RoomService) CreateRoom(room models.RoomCreate) (int, error) {
	return s.repo.CreateRoom(room)
}
