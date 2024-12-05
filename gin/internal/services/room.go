package services

import (
	"gin/internal/config"
	"gin/internal/models"
	"gin/internal/repositories"
)

type RoomService struct {
	repo *repositories.Repository
	cfg  *config.Config
}

func NewRoomService(repo *repositories.Repository, cfg *config.Config) *RoomService {
	return &RoomService{
		repo: repo,
		cfg:  cfg,
	}
}

func (s *RoomService) GetRoomById(id int) (models.Room, error) {
	return s.repo.GetRoomById(id)
}

func (s *RoomService) CreateRoom(room models.RoomCreate) (int, error) {
	return s.repo.CreateRoom(room)
}
