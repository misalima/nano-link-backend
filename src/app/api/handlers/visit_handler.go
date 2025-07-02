package handlers

import (
	"github.com/misalima/nano-link-backend/src/core/domain"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/misalima/nano-link-backend/src/app/api/handlers/dto"
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
	resp := struct {
		Count int `json:"count"`
	}{
		Count: len(visits),
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *VisitHandler) GetVisitHistory(c echo.Context) error {
	shortID := c.Param("short_id")
	if shortID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "short_id is required"})
	}
	url, err := h.urlService.GetURLByShortID(c.Request().Context(), shortID)
	if err == domain.ErrURLNotFound {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "URL not found"})
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	visits, err := h.urlService.GetVisitHistory(c.Request().Context(), url.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	var resp []dto.URLVisitResponse
	for _, v := range visits {
		resp = append(resp, dto.URLVisitResponse{
			ID:        v.ID.String(),
			URLID:     v.URLID.String(),
			VisitedAt: v.VisitedAt.Format("2006-01-02T15:04:05Z07:00"),
		})
	}
	return c.JSON(http.StatusOK, resp)
}
