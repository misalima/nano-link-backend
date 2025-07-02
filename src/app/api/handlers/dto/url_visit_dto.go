package dto

type CreateURLVisitRequest struct {
	URLID string `json:"url_id" binding:"required"`
}

type URLVisitResponse struct {
	ID        string `json:"id"`
	URLID     string `json:"url_id"`
	VisitedAt string `json:"visited_at"`
}
