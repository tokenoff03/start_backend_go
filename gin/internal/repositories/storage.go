package repositories

import (
	"fmt"
	"gin/internal/models"
)

type Storage struct {
	Rooms map[int]models.Room
}

func NewStorage() *Storage {
	rooms := make(map[int]models.Room)
	rooms[1] = models.Room{1, "100", "Single", "ASDASDASDAS"}
	return &Storage{
		Rooms: rooms,
	}
}

func (s *Storage) GetRoomById(id int) (models.Room, error) {
	//Если нет ключа в мап, то будет возвращен нулевое значение. Это у нас пустая структура Room
	return s.Rooms[id], nil
}

func (s *Storage) CreateRoom(room models.Room) (int, string) { //пока возвращаем стринг так как мы сами создаем ошибку и отправляем ввиду строки
	if _, exist := s.Rooms[room.Id]; exist == true {
		err := fmt.Sprint("Can not create room, room already exist")
		return 0, err
	}

	s.Rooms[room.Id] = room
	return room.Id, ""
}

func (s *Storage) GetAllRooms() []models.Room {
	var rooms []models.Room
	for _, room := range s.Rooms {
		rooms = append(rooms, room)
	}

	return rooms
}

func (s *Storage) DeleteRoom(id int) (int, string) {
	if _, exist := s.Rooms[id]; exist == true {
		err := fmt.Sprint("Can not delete room, room does not exist")
		return 0, err
	}

	delete(s.Rooms, id)
	return id, ""
}
