package models

import "time"

type Todo struct {
	ID              int       `json:"id" db:"id"`
	Title           string    `json:"title" db:"title"`
	TodoDescription string    `json:"todo_description" db:"todo_description"`
	Completed       bool      `json:"completed" db:"completed"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

