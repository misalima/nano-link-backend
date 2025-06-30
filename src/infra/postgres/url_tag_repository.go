package postgres

import (
	"context"
	"github.com/google/uuid"

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

func (r *URLTagRepository) Save(ctx context.Context, urlTag *domain.URLTag) error {
	panic("unimplemented")
}

func (r *URLTagRepository) FetchByURLID(ctx context.Context, urlID uuid.UUID) ([]*domain.URLTag, error) {
	panic("unimplemented")
}

func (r *URLTagRepository) FetchByTagID(ctx context.Context, tagID uuid.UUID) ([]*domain.URLTag, error) {
	panic("unimplemented")
}

func (r *URLTagRepository) Delete(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}

func (r *URLTagRepository) DeleteByURLAndTag(ctx context.Context, urlID, tagID uuid.UUID) error {
	panic("unimplemented")
}
