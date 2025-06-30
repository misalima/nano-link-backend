package domain

import (
	"errors"
	"github.com/google/uuid"
	"net/url"
	"time"
)

type URL struct {
	ID            uuid.UUID `json:"id"`
	ShortID       string    `json:"short_id"`
	CustomShortID *string   `json:"custom_short_id,omitempty"`
	OriginalURL   string    `json:"original_url"`
	TotalVisits   int       `json:"total_visits"`
	UserID        uuid.UUID `json:"user_id,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
}

func NewURL(shortID, originalURL string, userID uuid.UUID) *URL {
	return &URL{
		ID:          uuid.New(),
		ShortID:     shortID,
		OriginalURL: originalURL,
		UserID:      userID,
		TotalVisits: 0,
		CreatedAt:   time.Now(),
	}
}

func NewCustomURL(shortID, customShortID, originalURL string, userID uuid.UUID) *URL {
	return &URL{
		ID:            uuid.New(),
		ShortID:       shortID,
		CustomShortID: &customShortID,
		OriginalURL:   originalURL,
		UserID:        userID,
		TotalVisits:   0,
		CreatedAt:     time.Now(),
	}
}

func (u *URL) IncrementVisits() {
	u.TotalVisits++
}

func ValidateURL(originalURL string) error {
	if originalURL == "" {
		return errors.New("URL is empty")
	}

	parsedURL, err := url.ParseRequestURI(originalURL)
	if err != nil {
		return errors.New("URL is not valid")
	}

	if parsedURL.Scheme != "https" && parsedURL.Scheme != "http" {
		return errors.New("URL must start with http:// or https://")
	}

	if parsedURL.Host == "" {
		return errors.New("URL must have a valid host")
	}

	return nil
}
