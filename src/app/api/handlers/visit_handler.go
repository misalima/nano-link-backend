package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/misalima/nano-link-backend/src/core/ports"
)

type VisitHandler struct {
	urlService ports.URLService
}

func NewVisitHandler(urlService ports.URLService) *VisitHandler {
	return &VisitHandler{
		urlService: urlService,
	}
}

func (h *VisitHandler) GetVisitCount(c echo.Context) error {
	shortID := c.Param("short_id")
	if shortID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "short_id is required"})
	}
	url, err := h.urlService.GetURLByShortID(c.Request().Context(), shortID)
	if err != nil || url == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "URL not found"})
	}
	visits, err := h.urlService.GetVisitHistory(c.Request().Context(), url.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]int{"count": len(visits)})
}

func (h *VisitHandler) GetVisitHistory(c echo.Context) error {
	shortID := c.Param("short_id")
	if shortID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "short_id is required"})
	}
	url, err := h.urlService.GetURLByShortID(c.Request().Context(), shortID)
	if err != nil || url == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "URL not found"})
	}
	visits, err := h.urlService.GetVisitHistory(c.Request().Context(), url.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, visits)
}
