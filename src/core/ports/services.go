package ports

import (
	"context"

	"github.com/misalima/nano-link-backend/src/core/domain"
)

type UserService interface {
	Register(ctx context.Context, username, email, password string) (*domain.User, error)
	Authenticate(ctx context.Context, usernameOrEmail, password string) (*domain.User, error)
	GetUserByID(ctx context.Context, id string) (*domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User) error
	DeleteUser(ctx context.Context, id string) error
}

type URLService interface {
	CreateShortURL(ctx context.Context, originalURL string, userID string) (*domain.URL, error)
	CreateCustomShortURL(ctx context.Context, originalURL, customShortID string, userID string) (*domain.URL, error)
	GetURLByShortID(ctx context.Context, shortID string) (*domain.URL, error)
	GetURLByCustomShortID(ctx context.Context, customShortID string) (*domain.URL, error)
	GetURLsByUserID(ctx context.Context, userID string) ([]*domain.URL, error)
	UpdateURL(ctx context.Context, url *domain.URL) error
	DeleteURL(ctx context.Context, id string, userID string) error
	RecordVisit(ctx context.Context, urlID string) error
}

type TagService interface {
	CreateTag(ctx context.Context, name string) (*domain.Tag, error)
	GetTagByID(ctx context.Context, id string) (*domain.Tag, error)
	GetTagByName(ctx context.Context, name string) (*domain.Tag, error)
	GetAllTags(ctx context.Context) ([]*domain.Tag, error)
	UpdateTag(ctx context.Context, tag *domain.Tag) error
	DeleteTag(ctx context.Context, id string) error
}

type URLTagService interface {
	AddTagToURL(ctx context.Context, urlID, tagID string) error
	RemoveTagFromURL(ctx context.Context, urlID, tagID string) error
	GetTagsByURLID(ctx context.Context, urlID string) ([]*domain.Tag, error)
	GetURLsByTagID(ctx context.Context, tagID string) ([]*domain.URL, error)
}
