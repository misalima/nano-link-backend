package domain

import (
	"github.com/google/uuid"
	"time"
)

type URLVisit struct {
	ID        uuid.UUID `json:"id"`
	URLID     uuid.UUID `json:"url_id"`
	VisitedAt time.Time `json:"visited_at"`
}

func NewURLVisit(urlID uuid.UUID) *URLVisit {
	return &URLVisit{
		ID:        uuid.New(),
		URLID:     urlID,
		VisitedAt: time.Now(),
	}
}
