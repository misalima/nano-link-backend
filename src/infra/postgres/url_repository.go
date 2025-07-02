package postgres

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/misalima/nano-link-backend/src/infra/logger"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/misalima/nano-link-backend/src/core/domain"
	"github.com/misalima/nano-link-backend/src/core/ports"
)

type URLRepository struct {
	db *pgxpool.Pool
}

func NewURLRepository(db *pgxpool.Pool) ports.URLRepository {
	return &URLRepository{
		db: db,
	}
}

func (r *URLRepository) Save(ctx context.Context, url *domain.URL) error {
	query := `INSERT INTO urls (id, short_id, custom_short_id, original_url, user_id) VALUES ($1, $2, $3, $4, $5)`

	_, err := r.db.Exec(ctx, query, url.ID, url.ShortID, url.CustomShortID, url.OriginalURL, url.UserID)
	if err != nil {
		logger.Errorf("failed to insert url: %v", err)
		return err
	}

	logger.Infof("URL saved with ID: %s", url.ID)

	return nil
}

func (r *URLRepository) FetchByID(ctx context.Context, id uuid.UUID) (*domain.URL, error) {
	query := `SELECT id, short_id, custom_short_id, original_url, total_visits, user_id, created_at FROM urls WHERE id = $1`

	row := r.db.QueryRow(ctx, query, id)

	var url domain.URL
	err := row.Scan(&url.ID, &url.ShortID, &url.CustomShortID, &url.OriginalURL, &url.TotalVisits, &url.UserID, &url.CreatedAt)
	if err != nil {
		logger.Errorf("failed to fetch url by id: %v", err)
		return nil, err
	}

	return &url, nil
}

func (r *URLRepository) FetchByShortID(ctx context.Context, shortID string) (*domain.URL, error) {
	query := `SELECT id, short_id, custom_short_id, original_url, total_visits, user_id, created_at FROM urls WHERE short_id = $1`

	row := r.db.QueryRow(ctx, query, shortID)

	var url domain.URL
	err := row.Scan(&url.ID, &url.ShortID, &url.CustomShortID, &url.OriginalURL, &url.TotalVisits, &url.UserID, &url.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Infof("URL not found for short ID: %s", shortID)
			return nil, domain.ErrURLNotFound
		}
		logger.Errorf("failed to fetch url by short id: %v", err)
		return nil, err
	}

	return &url, nil
}

func (r *URLRepository) FetchByCustomShortID(ctx context.Context, customShortID string) (*domain.URL, error) {
	query := `SELECT id, short_id, custom_short_id, original_url, total_visits, user_id, created_at FROM urls WHERE custom_short_id = $1`

	row := r.db.QueryRow(ctx, query, customShortID)

	var url domain.URL
	err := row.Scan(&url.ID, &url.ShortID, &url.CustomShortID, &url.OriginalURL, &url.TotalVisits, &url.UserID, &url.CreatedAt)
	if err != nil {
		logger.Errorf("failed to fetch url by custom short id: %v", err)
		return nil, err
	}

	return &url, nil
}

func (r *URLRepository) FetchByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.URL, error) {
	query := `SELECT id, short_id, custom_short_id, original_url, total_visits, user_id, created_at FROM urls WHERE user_id = $1`
	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		logger.Errorf("failed to fetch urls by user id: %v", err)
		return nil, err
	}
	defer rows.Close()

	var urls []*domain.URL
	for rows.Next() {
		var url domain.URL
		err := rows.Scan(&url.ID, &url.ShortID, &url.CustomShortID, &url.OriginalURL, &url.TotalVisits, &url.UserID, &url.CreatedAt)
		if err != nil {
			logger.Errorf("failed to scan url row: %v", err)
			return nil, err
		}
		urls = append(urls, &url)
	}

	if err := rows.Err(); err != nil {
		logger.Errorf("error iterating over rows: %v", err)
		return nil, err
	}

	return urls, nil
}

func (r *URLRepository) Update(ctx context.Context, url *domain.URL) error {
	query := `UPDATE urls SET short_id = $1, custom_short_id = $2, original_url = $3, total_visits = $4 WHERE id = $5`

	_, err := r.db.Exec(ctx, query, url.ShortID, url.CustomShortID, url.OriginalURL, url.TotalVisits, url.ID)
	if err != nil {
		logger.Errorf("failed to update url: %v", err)
		return err
	}

	logger.Infof("URL updated with ID: %s", url.ID)

	return nil
}

func (r *URLRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM urls WHERE id = $1`
	tags, err := r.db.Exec(ctx, query, id)
	if err != nil {
		logger.Errorf("failed to delete url: %v", err)
		return err
	}

	if tags.RowsAffected() == 0 {
		logger.Warnf("No URL found with ID: %s", id)
		return nil
	}
	logger.Infof("URL deleted with ID: %s", id)
	return nil
}
