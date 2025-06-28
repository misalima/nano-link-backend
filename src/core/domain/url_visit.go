package domain

import (
	"time"
)

type URLVisit struct {
	ID        string    `json:"id"`
	URLID     string    `json:"url_id"`
	VisitedAt time.Time `json:"visited_at"`
}

func NewURLVisit(urlID string) *URLVisit {
	return &URLVisit{
		URLID:     urlID,
		VisitedAt: time.Now(),
	}
}
