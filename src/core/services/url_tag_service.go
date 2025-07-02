package services

import (
	"context"
	"errors"
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

func (s *URLTagService) AddTagToURL(ctx context.Context, urlID uuid.UUID, tagName string) error {
	url, err := s.urlRepo.FetchByID(ctx, urlID)
	if err != nil {
		return err
	}
	if url == nil {
		return domain.ErrURLNotFound
	}

	var tagID uuid.UUID

	existingTag, err := s.tagRepo.FetchByName(ctx, tagName)
	if err != nil {
		if errors.Is(err, domain.ErrTagNotFound) {
			return err
		}
	}

	if existingTag == nil {
		newTag := domain.NewTag(tagName)
		err = s.tagRepo.Save(ctx, newTag)
		if err != nil {
			return err
		}
		tagID = newTag.ID
	} else {
		tagID = existingTag.ID
	}

	urlTagExists, err := s.urlTagRepo.Exists(ctx, urlID, tagID)
	if err != nil {
		return err
	}

	if urlTagExists {
		return nil
	}

	urlTag := domain.NewURLTag(urlID, tagID)
	return s.urlTagRepo.Save(ctx, urlTag)
}

func (s *URLTagService) RemoveTagFromURL(ctx context.Context, urlID uuid.UUID, tagName string) error {
	url, err := s.urlRepo.FetchByID(ctx, urlID)
	if err != nil {
		return err
	}
	if url == nil {
		return domain.ErrURLNotFound
	}

	tag, err := s.tagRepo.FetchByName(ctx, tagName)
	if err != nil {
		return err
	}
	if tag == nil {
		return domain.ErrTagNameNotFound
	}

	return s.urlTagRepo.DeleteByURLAndTag(ctx, urlID, tag.ID)
}

func (s *URLTagService) GetTagsByURLID(ctx context.Context, urlID uuid.UUID) ([]*domain.Tag, error) {
	url, err := s.urlRepo.FetchByID(ctx, urlID)
	if err != nil {
		return nil, err
	}
	if url == nil {
		return nil, domain.ErrURLNotFound
	}

	urlTags, err := s.urlTagRepo.FetchByURLID(ctx, urlID)
	if err != nil {
		return nil, err
	}

	var tags []*domain.Tag
	for _, urlTag := range urlTags {
		tag, err := s.tagRepo.FetchByID(ctx, urlTag.TagID)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}

	return tags, nil
}

func (s *URLTagService) GetURLsByTagID(ctx context.Context, tagID uuid.UUID) ([]*domain.URL, error) {
	tag, err := s.tagRepo.FetchByID(ctx, tagID)
	if err != nil {
		return nil, err
	}
	if tag == nil {
		return nil, domain.ErrTagNotFound
	}

	urlTags, err := s.urlTagRepo.FetchByTagID(ctx, tagID)
	if err != nil {
		return nil, err
	}

	var urls []*domain.URL
	for _, urlTag := range urlTags {
		url, err := s.urlRepo.FetchByID(ctx, urlTag.URLID)
		if err != nil {
			return nil, err
		}
		urls = append(urls, url)
	}

	return urls, nil
}
