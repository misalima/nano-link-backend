package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/misalima/nano-link-backend/src/core/domain"
	"github.com/misalima/nano-link-backend/src/core/ports"
)

type URLVisitRepository struct {
	db *pgxpool.Pool
}

func NewURLVisitRepository(db *pgxpool.Pool) ports.URLVisitRepository {
	return &URLVisitRepository{
		db: db,
	}
}

func (r *URLVisitRepository) Save(ctx context.Context, visit *domain.URLVisit) error {
	panic("unimplemented")
}

func (r *URLVisitRepository) FetchByURLID(ctx context.Context, urlID string) ([]*domain.URLVisit, error) {
	panic("unimplemented")
}

func (r *URLVisitRepository) FetchByID(ctx context.Context, id string) (*domain.URLVisit, error) {
	panic("unimplemented")
}
