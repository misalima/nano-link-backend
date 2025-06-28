package domain

import (
	"time"
)

type Tag struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func NewTag(name string) *Tag {
	return &Tag{
		Name:      name,
		CreatedAt: time.Now(),
	}
}
