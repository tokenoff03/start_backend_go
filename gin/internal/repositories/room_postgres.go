package repositories

import (
	"fmt"
	"gin/internal/models"

	"github.com/jmoiron/sqlx"
)

type RoomPostgres struct {
	db *sqlx.DB
}

func NewRoomPostgres(db *sqlx.DB) *RoomPostgres {
	return &RoomPostgres{db: db}
}

func (r *RoomPostgres) GetRoomById(id int) (models.Room, error) {
	var room models.Room

	query := fmt.Sprintf("SELECT r.id, r.number, r.type, r.description FROM %s r WHERE r.id=$1", "room")
	err := r.db.Get(&room, query, id)

	if err != nil {
		return room, err
	}

	return room, nil
}

func (r *RoomPostgres) CreateRoom(room models.RoomCreate) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (number,type,description) values ($1, $2, $3) RETURNING id", "room")
	row := r.db.QueryRow(query, room.Number, room.Type, room.Description)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
