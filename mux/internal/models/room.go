package models

type Room struct {
	Id          int    `json: "id"`
	Number      string `json: "number"`
	Type        string `json: "type"`
	Description string `json: "description"`
}
