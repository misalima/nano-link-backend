package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/misalima/nano-link-backend/src/core/ports"
)

type URLHandler struct {
	urlService ports.URLService
}

func NewURLHandler(urlService ports.URLService) *URLHandler {
	return &URLHandler{
		urlService: urlService,
	}
}

func (h *URLHandler) CreateShortURL(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Create short URL feature is not implemented yet",
	})
}

func (h *URLHandler) GetURLDetails(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Get URL details feature is not implemented yet",
	})
}

func (h *URLHandler) RedirectToOriginalURL(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Redirect to original URL feature is not implemented yet",
	})
}

func (h *URLHandler) DeleteURL(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Delete URL feature is not implemented yet",
	})
}
