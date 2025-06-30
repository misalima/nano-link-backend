package domain

import (
	"github.com/google/uuid"
	"time"
)

type Tag struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func NewTag(name string) *Tag {
	return &Tag{
		ID:        uuid.New(),
		Name:      name,
		CreatedAt: time.Now(),
	}
}
