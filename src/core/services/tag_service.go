package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/misalima/nano-link-backend/src/infra/logger"
	"strings"

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
	if name = strings.TrimSpace(name); name == "" {
		logger.Error("Tag name cannot be empty")
		return nil, domain.ErrInvalidInput
	}

	// Check if tag already exists
	existingTag, err := s.tagRepo.FetchByName(ctx, name)
	if err == nil && existingTag != nil {
		logger.Infof("Tag with name '%s' already exists", name)
		return existingTag, nil
	}

	tag := domain.NewTag(name)
	err = s.tagRepo.Save(ctx, tag)
	if err != nil {
		logger.Errorf("Error saving tag: %v", err)
		return nil, err
	}

	return tag, nil
}

func (s *TagService) GetTagByID(ctx context.Context, id uuid.UUID) (*domain.Tag, error) {
	if id == uuid.Nil {
		logger.Error("Tag ID cannot be nil")
		return nil, domain.ErrInvalidInput
	}

	tag, err := s.tagRepo.FetchByID(ctx, id)
	if err != nil {
		logger.Errorf("Error fetching tag by ID: %v", err)
		return nil, err
	}

	return tag, nil
}

func (s *TagService) GetTagByName(ctx context.Context, name string) (*domain.Tag, error) {
	if name = strings.TrimSpace(name); name == "" {
		logger.Error("Tag name cannot be empty")
		return nil, domain.ErrInvalidInput
	}

	tag, err := s.tagRepo.FetchByName(ctx, name)
	if err != nil {
		logger.Errorf("Error fetching tag by name: %v", err)
		return nil, err
	}

	return tag, nil
}

func (s *TagService) GetAllTags(ctx context.Context) ([]*domain.Tag, error) {
	tags, err := s.tagRepo.FetchAll(ctx)
	if err != nil {
		logger.Errorf("Error fetching all tags: %v", err)
		return nil, err
	}

	return tags, nil
}

func (s *TagService) UpdateTag(ctx context.Context, tag *domain.Tag) error {
	if tag == nil {
		logger.Error("Tag cannot be nil")
		return domain.ErrInvalidInput
	}

	if tag.ID == uuid.Nil {
		logger.Error("Tag ID cannot be nil")
		return domain.ErrInvalidInput
	}

	if tag.Name = strings.TrimSpace(tag.Name); tag.Name == "" {
		logger.Error("Tag name cannot be empty")
		return domain.ErrInvalidInput
	}

	// Check if tag exists
	existingTag, err := s.tagRepo.FetchByID(ctx, tag.ID)
	if err != nil || existingTag == nil {
		logger.Errorf("Tag with ID %s not found", tag.ID)
		return domain.ErrNotFound
	}

	// Check if new name already exists for another tag
	if existingTag.Name != tag.Name {
		tagWithSameName, err := s.tagRepo.FetchByName(ctx, tag.Name)
		if err == nil && tagWithSameName != nil && tagWithSameName.ID != tag.ID {
			logger.Errorf("Tag with name '%s' already exists", tag.Name)
			return domain.ErrDuplicateEntry
		}
	}

	err = s.tagRepo.Update(ctx, tag)
	if err != nil {
		logger.Errorf("Error updating tag: %v", err)
		return err
	}

	return nil
}

func (s *TagService) DeleteTag(ctx context.Context, id uuid.UUID) error {
	if id == uuid.Nil {
		logger.Error("Tag ID cannot be nil")
		return domain.ErrInvalidInput
	}

	// Check if tag exists
	existingTag, err := s.tagRepo.FetchByID(ctx, id)
	if err != nil || existingTag == nil {
		logger.Errorf("Tag with ID %s not found", id)
		return domain.ErrNotFound
	}

	err = s.tagRepo.Delete(ctx, id)
	if err != nil {
		logger.Errorf("Error deleting tag: %v", err)
		return err
	}

	return nil
}
