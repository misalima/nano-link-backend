package domain

import (
	"github.com/google/uuid"
	"time"
)

type URLTag struct {
	ID        uuid.UUID `json:"id"`
	URLID     uuid.UUID `json:"url_id"`
	TagID     uuid.UUID `json:"tag_id"`
	CreatedAt time.Time `json:"created_at"`
}

func NewURLTag(urlID, tagID uuid.UUID) *URLTag {
	return &URLTag{
		ID:        uuid.New(),
		URLID:     urlID,
		TagID:     tagID,
		CreatedAt: time.Now(),
	}
}
