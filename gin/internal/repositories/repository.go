package repositories

import (
	"gin/internal/models"

	"github.com/jmoiron/sqlx"
)

type Room interface {
	GetRoomById(id int) (models.Room, error)
	CreateRoom(room models.RoomCreate) (int, error)
}

type Repository struct {
	Room
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Room: NewRoomPostgres(db),
	}
}
