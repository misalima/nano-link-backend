package ports

import (
	"context"
	"github.com/google/uuid"

	"github.com/misalima/nano-link-backend/src/core/domain"
)

type UserRepository interface {
	Save(ctx context.Context, user *domain.User) error
	FetchByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
	FetchByUsername(ctx context.Context, username string) (*domain.User, error)
	FetchByEmail(ctx context.Context, email string) (*domain.User, error)
	Update(ctx context.Context, user *domain.User) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type URLRepository interface {
	Save(ctx context.Context, url *domain.URL) error
	FetchByID(ctx context.Context, id uuid.UUID) (*domain.URL, error)
	FetchByShortID(ctx context.Context, shortID string) (*domain.URL, error)
	FetchByCustomShortID(ctx context.Context, customShortID string) (*domain.URL, error)
	FetchByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.URL, error)
	Update(ctx context.Context, url *domain.URL) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type TagRepository interface {
	Save(ctx context.Context, tag *domain.Tag) error
	FetchByID(ctx context.Context, id uuid.UUID) (*domain.Tag, error)
	FetchByName(ctx context.Context, name string) (*domain.Tag, error)
	FetchAll(ctx context.Context) ([]*domain.Tag, error)
	Update(ctx context.Context, tag *domain.Tag) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type URLTagRepository interface {
	Save(ctx context.Context, urlTag *domain.URLTag) error
	FetchByURLID(ctx context.Context, urlID uuid.UUID) ([]*domain.URLTag, error)
	FetchByTagID(ctx context.Context, tagID uuid.UUID) ([]*domain.URLTag, error)
	Delete(ctx context.Context, id uuid.UUID) error
	DeleteByURLAndTag(ctx context.Context, urlID, tagID uuid.UUID) error
	Exists(ctx context.Context, urlID, tagID uuid.UUID) (bool, error)
}

type URLVisitRepository interface {
	Save(ctx context.Context, visit *domain.URLVisit) error
	FetchByURLID(ctx context.Context, urlID uuid.UUID) ([]*domain.URLVisit, error)
	FetchByID(ctx context.Context, id uuid.UUID) (*domain.URLVisit, error)
}
