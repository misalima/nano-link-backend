package services

import (
	"context"
	"github.com/google/uuid"

	"github.com/misalima/nano-link-backend/src/core/domain"
	"github.com/misalima/nano-link-backend/src/core/ports"
)

type URLTagService struct {
	urlTagRepo ports.URLTagRepository
	urlRepo    ports.URLRepository
	tagRepo    ports.TagRepository
}

func NewURLTagService(urlTagRepo ports.URLTagRepository, urlRepo ports.URLRepository, tagRepo ports.TagRepository) ports.URLTagService {
	return &URLTagService{
		urlTagRepo: urlTagRepo,
		urlRepo:    urlRepo,
		tagRepo:    tagRepo,
	}
}

func (s *URLTagService) AddTagToURL(ctx context.Context, urlID, tagID uuid.UUID) error {
	panic("unimplemented")
}

func (s *URLTagService) RemoveTagFromURL(ctx context.Context, urlID, tagID uuid.UUID) error {
	panic("unimplemented")
}

func (s *URLTagService) GetTagsByURLID(ctx context.Context, urlID uuid.UUID) ([]*domain.Tag, error) {
	panic("unimplemented")
}

func (s *URLTagService) GetURLsByTagID(ctx context.Context, tagID uuid.UUID) ([]*domain.URL, error) {
	panic("unimplemented")
}