package model

import "time"

type Todo struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	IsCompleted bool      `json:"isCompleted"`
	DueDate     time.Time `json:"dueDate,omitempty"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
}
