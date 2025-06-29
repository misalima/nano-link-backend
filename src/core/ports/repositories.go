package ports

import (
	"context"

	"github.com/misalima/nano-link-backend/src/core/domain"
)

type UserRepository interface {
	Save(ctx context.Context, user *domain.User) error
	FetchByID(ctx context.Context, id string) (*domain.User, error)
	FetchByUsername(ctx context.Context, username string) (*domain.User, error)
	FetchByEmail(ctx context.Context, email string) (*domain.User, error)
	Update(ctx context.Context, user *domain.User) error
	Delete(ctx context.Context, id string) error
}

type URLRepository interface {
	Save(ctx context.Context, url *domain.URL) error
	FetchByID(ctx context.Context, id string) (*domain.URL, error)
	FetchByShortID(ctx context.Context, shortID string) (*domain.URL, error)
	FetchByCustomShortID(ctx context.Context, customShortID string) (*domain.URL, error)
	FetchByUserID(ctx context.Context, userID string) ([]*domain.URL, error)
	Update(ctx context.Context, url *domain.URL) error
	Delete(ctx context.Context, id string) error
}

type TagRepository interface {
	Save(ctx context.Context, tag *domain.Tag) error
	FetchByID(ctx context.Context, id string) (*domain.Tag, error)
	FetchByName(ctx context.Context, name string) (*domain.Tag, error)
	FetchAll(ctx context.Context) ([]*domain.Tag, error)
	Update(ctx context.Context, tag *domain.Tag) error
	Delete(ctx context.Context, id string) error
}

type URLTagRepository interface {
	Save(ctx context.Context, urlTag *domain.URLTag) error
	FetchByURLID(ctx context.Context, urlID string) ([]*domain.URLTag, error)
	FetchByTagID(ctx context.Context, tagID string) ([]*domain.URLTag, error)
	Delete(ctx context.Context, id string) error
	DeleteByURLAndTag(ctx context.Context, urlID, tagID string) error
}

type URLVisitRepository interface {
	Save(ctx context.Context, visit *domain.URLVisit) error
	FetchByURLID(ctx context.Context, urlID string) ([]*domain.URLVisit, error)
	FetchByID(ctx context.Context, id string) (*domain.URLVisit, error)
}
