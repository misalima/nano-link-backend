package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/misalima/nano-link-backend/src/infra/logger"
	"github.com/misalima/nano-link-backend/src/utils"

	"github.com/misalima/nano-link-backend/src/core/domain"
	"github.com/misalima/nano-link-backend/src/core/ports"
)

type URLService struct {
	urlRepo      ports.URLRepository
	urlVisitRepo ports.URLVisitRepository
}

func NewURLService(urlRepo ports.URLRepository, urlVisitRepo ports.URLVisitRepository) ports.URLService {
	return &URLService{
		urlRepo:      urlRepo,
		urlVisitRepo: urlVisitRepo,
	}
}

func (s *URLService) CreateShortURL(ctx context.Context, originalURL string, userID uuid.UUID) (*domain.URL, error) {
	err := domain.ValidateURL(originalURL)
	if err != nil {
		logger.Errorf("Error validating URL: %v", err)
		return nil, err
	}

	shortID, err := utils.HashURLWithRandom(originalURL)
	if err != nil {
		logger.Errorf("Error creating short ID from URL: %v", err)
		return nil, err
	}

	url := domain.NewURL(shortID, originalURL, userID)
	url.CustomShortID = nil

	err = s.urlRepo.Save(ctx, url)
	if err != nil {
		logger.Errorf("Error saving URL: %v", err)
		return nil, err
	}

	return url, nil
}

func (s *URLService) CreateCustomShortURL(ctx context.Context, originalURL, customShortID string, userID uuid.UUID) (*domain.URL, error) {
	err := domain.ValidateURL(originalURL)
	if err != nil {
		logger.Errorf("Error validating URL: %v", err)
		return nil, err
	}

	if customShortID == "" {
		logger.Error("Custom short ID cannot be empty")
		return nil, domain.ErrInvalidCustomShortID
	}

	url := domain.NewURL(customShortID, originalURL, userID)
	url.CustomShortID = &customShortID
	url.ShortID = customShortID

	err = s.urlRepo.Save(ctx, url)
	if err != nil {
		logger.Errorf("Error saving URL with custom short ID: %v", err)
		return nil, err
	}

	return url, nil
}

func (s *URLService) GetURLByShortID(ctx context.Context, shortID string) (*domain.URL, error) {
	url, err := s.urlRepo.FetchByShortID(ctx, shortID)
	if err != nil {
		logger.Errorf("failed to get url by short id: %v", err)
		return nil, err
	}

	return url, nil
}

func (s *URLService) GetURLByCustomShortID(ctx context.Context, customShortID string) (*domain.URL, error) {
	url, err := s.urlRepo.FetchByCustomShortID(ctx, customShortID)
	if err != nil {
		return nil, err
	}

	return url, nil
}

func (s *URLService) GetURLsByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.URL, error) {
	urls, err := s.urlRepo.FetchByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return urls, nil
}

func (s *URLService) UpdateURL(ctx context.Context, url *domain.URL) error {
	if url == nil {
		logger.Error("URL cannot be nil")
		return domain.ErrInvalidURL
	}

	err := domain.ValidateURL(url.OriginalURL)
	if err != nil {
		logger.Errorf("Error validating URL: %v", err)
		return err
	}

	if url.CustomShortID != nil && *url.CustomShortID == "" {
		logger.Error("Custom short ID cannot be empty")
		return domain.ErrInvalidCustomShortID
	}

	err = s.urlRepo.Save(ctx, url)
	if err != nil {
		return err
	}

	return nil
}

func (s *URLService) DeleteURL(ctx context.Context, id, userID uuid.UUID) error {
	return s.urlRepo.Delete(ctx, id)
}

func (s *URLService) RecordVisit(ctx context.Context, urlID uuid.UUID) error {
	visit := domain.NewURLVisit(urlID)
	err := s.urlVisitRepo.Save(ctx, visit)
	if err != nil {
		return err
	}

	return nil
}

func CreateShortIDFromURL(originalURL string) string {
	shortID, err := utils.HashURLWithRandom(originalURL)
	if err != nil {
		logger.Errorf("Error creating short ID from URL: %v", err)
		return ""
	}
	return shortID
}
