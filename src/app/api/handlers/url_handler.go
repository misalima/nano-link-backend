package handlers

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/misalima/nano-link-backend/src/app/api/auth"
	"github.com/misalima/nano-link-backend/src/app/api/handlers/dto"
	"github.com/misalima/nano-link-backend/src/core/domain"
	"github.com/misalima/nano-link-backend/src/core/ports"
	"github.com/misalima/nano-link-backend/src/infra/logger"
	"github.com/misalima/nano-link-backend/src/utils"
	"net/http"
)

type URLHandler struct {
	urlService ports.URLService
}

func NewURLHandler(urlService ports.URLService) *URLHandler {
	return &URLHandler{
		urlService: urlService,
	}
}

type CreateShortURLRequest = dto.CreateURLRequest

func (h *URLHandler) CreateShortURL(c echo.Context) error {
	var req dto.CreateURLRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	if req.OriginalURL == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Original URL is required",
		})
	}
	if req.CustomShortID != nil && *req.CustomShortID != "" {
		if len(*req.CustomShortID) < 3 || len(*req.CustomShortID) > 20 {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Custom short ID must be between 3 and 20 characters",
			})
		}
		if !utils.IsValidCustomShortID(*req.CustomShortID) {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Custom short ID can only contain alphanumeric characters and hyphens",
			})
		}
	}

	// Chama o serviço e monta a resposta usando o DTO de response
	var url *domain.URL
	var err error
	if req.CustomShortID != nil && *req.CustomShortID != "" {
		url, err = h.urlService.CreateCustomShortURL(c.Request().Context(), req.OriginalURL, *req.CustomShortID, uuid.Nil)
	} else {
		url, err = h.urlService.CreateShortURL(c.Request().Context(), req.OriginalURL, uuid.Nil)
	}
	if err != nil || url == nil {
		if err == domain.ErrInvalidURL {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid URL format",
			})
		}
		if err == domain.ErrInvalidCustomShortID {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid custom short ID",
			})
		}
		if err == domain.ErrCustomShortIDExists {
			return c.JSON(http.StatusConflict, map[string]string{
				"error": "Custom short ID already exists",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create short URL",
		})
	}
	resp := dto.URLResponse{
		ID:            url.ID.String(),
		ShortID:       url.ShortID,
		CustomShortID: url.CustomShortID,
		OriginalURL:   url.OriginalURL,
		TotalVisits:   url.TotalVisits,
		UserID:        url.UserID.String(),
		CreatedAt:     url.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
	return c.JSON(http.StatusCreated, resp)
}

func (h *URLHandler) GetURLDetails(c echo.Context) error {
	shortID := c.Param("short_id")
	if shortID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Short ID is required",
		})
	}

	url, err := h.urlService.GetURLByShortID(c.Request().Context(), shortID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to get URL details",
		})
	}
	if url == nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "URL not found",
		})
	}

	return c.JSON(http.StatusOK, url)
}

func (h *URLHandler) RedirectToOriginalURL(c echo.Context) error {
	shortId := c.Param("short_id")
	if shortId == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Short ID is required",
		})
	}

	url, err := h.urlService.GetURLByShortID(c.Request().Context(), shortId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to retrieve URL",
		})
	}

	if url == nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "URL not found",
		})
	}

	err = h.urlService.RecordVisit(c.Request().Context(), url.ID)
	if err != nil {
		logger.Warnf("Failed to record visit for URL %s: %v", url.ID, err)
	}

	return c.Redirect(http.StatusFound, url.OriginalURL)
}

func (h *URLHandler) DeleteURL(c echo.Context) error {
	idStr := c.Param("id")
	if idStr == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "URL ID is required",
		})
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid URL ID format",
		})
	}

	userID, err := auth.GetUserIDFromToken(c)
	if err != nil || userID == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "User not authenticated",
		})
	}

	err = h.urlService.DeleteURL(c.Request().Context(), id, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to delete URL",
		})
	}

	return c.NoContent(http.StatusNoContent)
}
