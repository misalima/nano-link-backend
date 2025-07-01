package handlers

import (
	"net/http"

	"github.com/google/uuid"
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
	urlIDStr := c.Param("short_id")
	tagIDStr := c.QueryParam("tag_id")
	if urlIDStr == "" || tagIDStr == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "short_id and tag_id are required"})
	}
	urlID, err := uuid.Parse(urlIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid short_id"})
	}
	tagID, err := uuid.Parse(tagIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid tag_id"})
	}
	err = h.urlTagService.AddTagToURL(c.Request().Context(), urlID, tagID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Tag added to URL"})
}

func (h *TagHandler) RemoveTagFromURL(c echo.Context) error {
	urlIDStr := c.Param("short_id")
	tagIDStr := c.Param("tag_id")
	if urlIDStr == "" || tagIDStr == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "short_id and tag_id are required"})
	}
	urlID, err := uuid.Parse(urlIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid short_id"})
	}
	tagID, err := uuid.Parse(tagIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid tag_id"})
	}
	err = h.urlTagService.RemoveTagFromURL(c.Request().Context(), urlID, tagID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Tag removed from URL"})
}

func (h *TagHandler) GetTagsForURL(c echo.Context) error {
	urlIDStr := c.Param("short_id")
	if urlIDStr == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "short_id is required"})
	}
	urlID, err := uuid.Parse(urlIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid short_id"})
	}
	tags, err := h.urlTagService.GetTagsByURLID(c.Request().Context(), urlID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, tags)
}
