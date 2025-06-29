package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/misalima/nano-link-backend/src/core/ports"
)

type TagHandler struct {
	tagService    ports.TagService
	urlTagService ports.URLTagService
}

func NewTagHandler(tagService ports.TagService, urlTagService ports.URLTagService) *TagHandler {
	return &TagHandler{
		tagService:    tagService,
		urlTagService: urlTagService,
	}
}

func (h *TagHandler) AddTagToURL(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Add tag to URL feature is not implemented yet",
	})
}

func (h *TagHandler) RemoveTagFromURL(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Remove tag from URL feature is not implemented yet",
	})
}

func (h *TagHandler) GetTagsForURL(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Get tags for URL feature is not implemented yet",
	})
}
