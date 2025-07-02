package dto

type CreateTagRequest struct {
	Name string `json:"name" binding:"required"`
}

type TagResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}
