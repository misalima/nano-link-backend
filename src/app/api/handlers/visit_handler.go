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
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Get visit count feature is not implemented yet",
	})
}

func (h *VisitHandler) GetVisitHistory(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Get visit history feature is not implemented yet",
	})
}
