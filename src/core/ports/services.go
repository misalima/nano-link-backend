package ports

import (
	"context"
	"github.com/google/uuid"

	"github.com/misalima/nano-link-backend/src/core/domain"
)

type UserService interface {
	Register(ctx context.Context, username, email, password string) (*domain.User, error)
	Authenticate(ctx context.Context, usernameOrEmail, password string) (*domain.User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

type URLService interface {
	CreateShortURL(ctx context.Context, originalURL string, userID uuid.UUID) (*domain.URL, error)
	CreateCustomShortURL(ctx context.Context, originalURL, customShortID string, userID uuid.UUID) (*domain.URL, error)
	GetURLByShortID(ctx context.Context, shortID string) (*domain.URL, error)
	GetURLByCustomShortID(ctx context.Context, customShortID string) (*domain.URL, error)
	GetURLsByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.URL, error)
	UpdateURL(ctx context.Context, url *domain.URL) error
	DeleteURL(ctx context.Context, id, userID uuid.UUID) error
	RecordVisit(ctx context.Context, urlID uuid.UUID) error
	GetVisitHistory(ctx context.Context, urlID uuid.UUID) ([]*domain.URLVisit, error)
}

type TagService interface {
	CreateTag(ctx context.Context, name string) (*domain.Tag, error)
	GetTagByID(ctx context.Context, id uuid.UUID) (*domain.Tag, error)
	GetTagByName(ctx context.Context, name string) (*domain.Tag, error)
	GetAllTags(ctx context.Context) ([]*domain.Tag, error)
	UpdateTag(ctx context.Context, tag *domain.Tag) error
	DeleteTag(ctx context.Context, id uuid.UUID) error
}

type URLTagService interface {
	AddTagToURL(ctx context.Context, urlID uuid.UUID, tagName string) error
	RemoveTagFromURL(ctx context.Context, urlID uuid.UUID, tagName string) error
	GetTagsByURLID(ctx context.Context, urlID uuid.UUID) ([]*domain.Tag, error)
	GetURLsByTagID(ctx context.Context, tagID uuid.UUID) ([]*domain.URL, error)
}
