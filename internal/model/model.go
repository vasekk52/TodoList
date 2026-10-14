package model

import "time"

type Todo struct {
	ID           int
	Title        string
	Completed    bool
	CreatedAt	 time.Time
	CompletedAt *time.Time
}

type Config struct {
	DB_HOST 	string
	DB_PORT 	string
	DB_USER 	string
	DB_PASS 	string
	DB_NAME 	string
	ServerPort	string
}

type TodoDTO struct {
	Title 		string
}