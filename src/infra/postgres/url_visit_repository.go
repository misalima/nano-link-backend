package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/misalima/nano-link-backend/src/infra/logger"

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
	query := `INSERT INTO url_visits (id, url_id, visited_at) VALUES ($1, $2, $3)`

	_, err := r.db.Exec(ctx, query, visit.ID, visit.URLID, visit.VisitedAt)
	if err != nil {
		logger.Errorf("failed to insert url visit: %v", err)
		return err
	}

	return nil
}

func (r *URLVisitRepository) FetchByURLID(ctx context.Context, urlID uuid.UUID) ([]*domain.URLVisit, error) {
	panic("unimplemented")
}

func (r *URLVisitRepository) FetchByID(ctx context.Context, id uuid.UUID) (*domain.URLVisit, error) {
	panic("unimplemented")
}
