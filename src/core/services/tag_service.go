package services

import (
	"context"
	"github.com/google/uuid"

	"github.com/misalima/nano-link-backend/src/core/domain"
	"github.com/misalima/nano-link-backend/src/core/ports"
)

type TagService struct {
	tagRepo ports.TagRepository
}

func NewTagService(tagRepo ports.TagRepository) ports.TagService {
	return &TagService{
		tagRepo: tagRepo,
	}
}

func (s *TagService) CreateTag(ctx context.Context, name string) (*domain.Tag, error) {
	panic("unimplemented")
}

func (s *TagService) GetTagByID(ctx context.Context, id uuid.UUID) (*domain.Tag, error) {
	panic("unimplemented")
}

func (s *TagService) GetTagByName(ctx context.Context, name string) (*domain.Tag, error) {
	panic("unimplemented")
}

func (s *TagService) GetAllTags(ctx context.Context) ([]*domain.Tag, error) {
	panic("unimplemented")
}

func (s *TagService) UpdateTag(ctx context.Context, tag *domain.Tag) error {
	panic("unimplemented")
}

func (s *TagService) DeleteTag(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}