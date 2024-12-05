package models

type Room struct {
	Id          int    `json:"id" db:"id"`
	Number      string `json:"number" db:"number"`
	Type        string `json:"type" db:"type"`
	Description string `json:"description" db:"description"`
}

type RoomCreate struct {
	Number      string `json:"number"`
	Type        string `json:"type"`
	Description string `json:"description"`
}
