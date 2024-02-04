package model

import "github.com/google/uuid"


type Task struct {
	ID uuid.UUID `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	IsDone bool `json:"is_done"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}