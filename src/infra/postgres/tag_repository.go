package postgres

import (
	"context"

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

func (r *TagRepository) Create(ctx context.Context, tag *domain.Tag) error {
	panic("unimplemented")
}

func (r *TagRepository) GetByID(ctx context.Context, id string) (*domain.Tag, error) {
	panic("unimplemented")
}

func (r *TagRepository) GetByName(ctx context.Context, name string) (*domain.Tag, error) {
	panic("unimplemented")
}

func (r *TagRepository) GetAll(ctx context.Context) ([]*domain.Tag, error) {
	panic("unimplemented")
}

func (r *TagRepository) Update(ctx context.Context, tag *domain.Tag) error {
	panic("unimplemented")
}

func (r *TagRepository) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}