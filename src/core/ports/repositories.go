package ports

import (
	"context"

	"github.com/misalima/nano-link-backend/src/core/domain"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	GetByID(ctx context.Context, id string) (*domain.User, error)
	GetByUsername(ctx context.Context, username string) (*domain.User, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	Update(ctx context.Context, user *domain.User) error
	Delete(ctx context.Context, id string) error
}

type URLRepository interface {
	Create(ctx context.Context, url *domain.URL) error
	GetByID(ctx context.Context, id string) (*domain.URL, error)
	GetByShortID(ctx context.Context, shortID string) (*domain.URL, error)
	GetByCustomShortID(ctx context.Context, customShortID string) (*domain.URL, error)
	GetByUserID(ctx context.Context, userID string) ([]*domain.URL, error)
	Update(ctx context.Context, url *domain.URL) error
	Delete(ctx context.Context, id string) error
}

type TagRepository interface {
	Create(ctx context.Context, tag *domain.Tag) error
	GetByID(ctx context.Context, id string) (*domain.Tag, error)
	GetByName(ctx context.Context, name string) (*domain.Tag, error)
	GetAll(ctx context.Context) ([]*domain.Tag, error)
	Update(ctx context.Context, tag *domain.Tag) error
	Delete(ctx context.Context, id string) error
}

type URLTagRepository interface {
	Create(ctx context.Context, urlTag *domain.URLTag) error
	GetByURLID(ctx context.Context, urlID string) ([]*domain.URLTag, error)
	GetByTagID(ctx context.Context, tagID string) ([]*domain.URLTag, error)
	Delete(ctx context.Context, id string) error
	DeleteByURLAndTag(ctx context.Context, urlID, tagID string) error
}

type URLVisitRepository interface {
	Create(ctx context.Context, visit *domain.URLVisit) error
	GetByURLID(ctx context.Context, urlID string) ([]*domain.URLVisit, error)
	GetByID(ctx context.Context, id string) (*domain.URLVisit, error)
}
