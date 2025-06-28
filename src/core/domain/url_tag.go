package domain

import (
	"time"
)

type URLTag struct {
	ID        string    `json:"id"`
	URLID     string    `json:"url_id"`
	TagID     string    `json:"tag_id"`
	CreatedAt time.Time `json:"created_at"`
}

func NewURLTag(urlID, tagID string) *URLTag {
	return &URLTag{
		URLID:     urlID,
		TagID:     tagID,
		CreatedAt: time.Now(),
	}
}
