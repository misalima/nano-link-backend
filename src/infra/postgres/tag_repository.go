package postgres

import (
	"context"
	"github.com/google/uuid"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/misalima/nano-link-backend/src/core/domain"
	"github.com/misalima/nano-link-backend/src/core/ports"
)

type TagRepository struct {
	db *pgxpool.Pool
}

func NewTagRepository(db *pgxpool.Pool) ports.TagRepository {
	return &TagRepository{
		db: db,
	}
}

func (r *TagRepository) Save(ctx context.Context, tag *domain.Tag) error {
	panic("unimplemented")
}

func (r *TagRepository) FetchByID(ctx context.Context, id uuid.UUID) (*domain.Tag, error) {
	panic("unimplemented")
}

func (r *TagRepository) FetchByName(ctx context.Context, name string) (*domain.Tag, error) {
	panic("unimplemented")
}

func (r *TagRepository) FetchAll(ctx context.Context) ([]*domain.Tag, error) {
	panic("unimplemented")
}

func (r *TagRepository) Update(ctx context.Context, tag *domain.Tag) error {
	panic("unimplemented")
}

func (r *TagRepository) Delete(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}
