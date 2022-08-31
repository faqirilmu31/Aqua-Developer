package model

type Customer struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Error struct {
	Message string `json:"message"`
}

var CustomerList = make(map[int]Customer)

var LastID int = 1