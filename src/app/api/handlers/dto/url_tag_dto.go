package dto

type CreateURLTagRequest struct {
	URLID   string `json:"url_id" binding:"required"`
	TagName string `json:"tag_name" binding:"required"`
}

type URLTagResponse struct {
	ID        string `json:"id"`
	URLID     string `json:"url_id"`
	TagID     string `json:"tag_id"`
	CreatedAt string `json:"created_at"`
}
