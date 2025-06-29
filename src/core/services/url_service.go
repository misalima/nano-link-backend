package services

import (
	"context"

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

func (s *URLService) CreateShortURL(ctx context.Context, originalURL string, userID string) (*domain.URL, error) {
	panic("unimplemented")
}

func (s *URLService) CreateCustomShortURL(ctx context.Context, originalURL, customShortID string, userID string) (*domain.URL, error) {
	panic("unimplemented")
}

func (s *URLService) GetURLByShortID(ctx context.Context, shortID string) (*domain.URL, error) {
	panic("unimplemented")
}

func (s *URLService) GetURLByCustomShortID(ctx context.Context, customShortID string) (*domain.URL, error) {
	panic("unimplemented")
}

func (s *URLService) GetURLsByUserID(ctx context.Context, userID string) ([]*domain.URL, error) {
	panic("unimplemented")
}

func (s *URLService) UpdateURL(ctx context.Context, url *domain.URL) error {
	panic("unimplemented")
}

func (s *URLService) DeleteURL(ctx context.Context, id string, userID string) error {
	panic("unimplemented")
}

func (s *URLService) RecordVisit(ctx context.Context, urlID string) error {
	panic("unimplemented")
}