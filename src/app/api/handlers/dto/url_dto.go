package dto

type CreateURLRequest struct {
	OriginalURL   string  `json:"original_url" binding:"required"`
	CustomShortID *string `json:"custom_short_id,omitempty"`
}

type URLResponse struct {
	ID            string  `json:"id"`
	ShortID       string  `json:"short_id"`
	CustomShortID *string `json:"custom_short_id,omitempty"`
	OriginalURL   string  `json:"original_url"`
	TotalVisits   int     `json:"total_visits"`
	UserID        string  `json:"user_id"`
	CreatedAt     string  `json:"created_at"`
}
