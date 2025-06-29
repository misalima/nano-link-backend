package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/misalima/nano-link-backend/src/core/domain"
	"github.com/misalima/nano-link-backend/src/core/ports"
)

type URLTagRepository struct {
	db *pgxpool.Pool
}

func NewURLTagRepository(db *pgxpool.Pool) ports.URLTagRepository {
	return &URLTagRepository{
		db: db,
	}
}

func (r *URLTagRepository) Create(ctx context.Context, urlTag *domain.URLTag) error {
	panic("unimplemented")
}

func (r *URLTagRepository) GetByURLID(ctx context.Context, urlID string) ([]*domain.URLTag, error) {
	panic("unimplemented")
}

func (r *URLTagRepository) GetByTagID(ctx context.Context, tagID string) ([]*domain.URLTag, error) {
	panic("unimplemented")
}

func (r *URLTagRepository) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

func (r *URLTagRepository) DeleteByURLAndTag(ctx context.Context, urlID, tagID string) error {
	panic("unimplemented")
}