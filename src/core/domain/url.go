package domain

import (
	"time"
)

type URL struct {
	ID            string    `json:"id"`
	ShortID       string    `json:"short_id"`
	CustomShortID string    `json:"custom_short_id,omitempty"`
	OriginalURL   string    `json:"original_url"`
	TotalVisits   int       `json:"total_visits"`
	UserID        string    `json:"user_id,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
}

func NewURL(shortID, originalURL string, userID string) *URL {
	return &URL{
		ShortID:     shortID,
		OriginalURL: originalURL,
		UserID:      userID,
		TotalVisits: 0,
		CreatedAt:   time.Now(),
	}
}

func NewCustomURL(shortID, customShortID, originalURL string, userID string) *URL {
	return &URL{
		ShortID:       shortID,
		CustomShortID: customShortID,
		OriginalURL:   originalURL,
		UserID:        userID,
		TotalVisits:   0,
		CreatedAt:     time.Now(),
	}
}

func (u *URL) IncrementVisits() {
	u.TotalVisits++
}
