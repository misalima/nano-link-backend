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
	query := `SELECT id, url_id, visited_at FROM url_visits WHERE url_id = $1`
	rows, err := r.db.Query(ctx, query, urlID)
	if err != nil {
		logger.Errorf("failed to fetch url visits by url_id: %v", err)
		return nil, err
	}
	defer rows.Close()

	var visits []*domain.URLVisit
	for rows.Next() {
		var visit domain.URLVisit
		err := rows.Scan(&visit.ID, &visit.URLID, &visit.VisitedAt)
		if err != nil {
			logger.Errorf("failed to scan url visit row: %v", err)
			return nil, err
		}
		visits = append(visits, &visit)
	}

	if err := rows.Err(); err != nil {
		logger.Errorf("error iterating over rows: %v", err)
		return nil, err
	}

	return visits, nil
}

func (r *URLVisitRepository) FetchByID(ctx context.Context, id uuid.UUID) (*domain.URLVisit, error) {
	query := `SELECT id, url_id, visited_at FROM url_visits WHERE id = $1`

	row := r.db.QueryRow(ctx, query, id)

	var visit domain.URLVisit
	err := row.Scan(&visit.ID, &visit.URLID, &visit.VisitedAt)
	if err != nil {
		logger.Errorf("failed to fetch url visit by id: %v", err)
		return nil, err
	}

	return &visit, nil
}
