package handlers

import (
	"errors"
	"github.com/misalima/nano-link-backend/src/core/domain"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/misalima/nano-link-backend/src/app/api/handlers/dto"
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
	var req dto.CreateURLTagRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}
	urlID, err := uuid.Parse(req.URLID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid url_id"})
	}

	if req.TagName == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Tag name is required"})
	}

	err = h.urlTagService.AddTagToURL(c.Request().Context(), urlID, req.TagName)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidURL) {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid URL ID"})
		} else if errors.Is(err, domain.ErrURLNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "URL not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Tag added to URL"})
}

func (h *TagHandler) RemoveTagFromURL(c echo.Context) error {
	urlIDStr := c.Param("url_id")
	tagName := c.Param("tag_name")
	if urlIDStr == "" || tagName == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "short_id and tag_id are required"})
	}
	urlID, err := uuid.Parse(urlIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid short_id"})
	}

	err = h.urlTagService.RemoveTagFromURL(c.Request().Context(), urlID, tagName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Tag removed from URL"})
}

func (h *TagHandler) GetTagsForURL(c echo.Context) error {
	urlIDStr := c.Param("url_id")
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

func (h *TagHandler) DeleteTag(c echo.Context) error {
	tagIDStr := c.Param("id")
	if tagIDStr == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Tag ID is required"})
	}
	tagID, err := uuid.Parse(tagIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid tag ID"})
	}
	err = h.tagService.DeleteTag(c.Request().Context(), tagID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Tag deleted successfully"})
}
