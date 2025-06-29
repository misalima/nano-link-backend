package services

import (
	"context"

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

func (s *URLTagService) AddTagToURL(ctx context.Context, urlID, tagID string) error {
	panic("unimplemented")
}

func (s *URLTagService) RemoveTagFromURL(ctx context.Context, urlID, tagID string) error {
	panic("unimplemented")
}

func (s *URLTagService) GetTagsByURLID(ctx context.Context, urlID string) ([]*domain.Tag, error) {
	panic("unimplemented")
}

func (s *URLTagService) GetURLsByTagID(ctx context.Context, tagID string) ([]*domain.URL, error) {
	panic("unimplemented")
}