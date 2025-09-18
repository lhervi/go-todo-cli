package main

import "time"

type Task struct {
	ID        int       `json:"id"` //json:id
	Content   string    `json:"content"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}
