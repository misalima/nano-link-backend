package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/misalima/nano-link-backend/src/infra/logger"

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
	query := `INSERT INTO url_tags (id, url_id, tag_id, created_at) VALUES ($1, $2, $3, $4)`

	_, err := r.db.Exec(ctx, query, urlTag.ID, urlTag.URLID, urlTag.TagID, urlTag.CreatedAt)
	if err != nil {
		logger.Errorf("failed to insert url_tag: %v", err)
		return err
	}

	logger.Infof("URL-Tag relationship saved with ID: %s", urlTag.ID)

	return nil
}

func (r *URLTagRepository) FetchByURLID(ctx context.Context, urlID uuid.UUID) ([]*domain.URLTag, error) {
	query := `SELECT id, url_id, tag_id, created_at FROM url_tags WHERE url_id = $1`
	rows, err := r.db.Query(ctx, query, urlID)
	if err != nil {
		logger.Errorf("failed to fetch url_tags by url_id: %v", err)
		return nil, err
	}
	defer rows.Close()

	var urlTags []*domain.URLTag
	for rows.Next() {
		var urlTag domain.URLTag
		err := rows.Scan(&urlTag.ID, &urlTag.URLID, &urlTag.TagID, &urlTag.CreatedAt)
		if err != nil {
			logger.Errorf("failed to scan url_tag row: %v", err)
			return nil, err
		}
		urlTags = append(urlTags, &urlTag)
	}

	if err := rows.Err(); err != nil {
		logger.Errorf("error iterating over rows: %v", err)
		return nil, err
	}

	return urlTags, nil
}

func (r *URLTagRepository) FetchByTagID(ctx context.Context, tagID uuid.UUID) ([]*domain.URLTag, error) {
	query := `SELECT id, url_id, tag_id, created_at FROM url_tags WHERE tag_id = $1`
	rows, err := r.db.Query(ctx, query, tagID)
	if err != nil {
		logger.Errorf("failed to fetch url_tags by tag_id: %v", err)
		return nil, err
	}
	defer rows.Close()

	var urlTags []*domain.URLTag
	for rows.Next() {
		var urlTag domain.URLTag
		err := rows.Scan(&urlTag.ID, &urlTag.URLID, &urlTag.TagID, &urlTag.CreatedAt)
		if err != nil {
			logger.Errorf("failed to scan url_tag row: %v", err)
			return nil, err
		}
		urlTags = append(urlTags, &urlTag)
	}

	if err := rows.Err(); err != nil {
		logger.Errorf("error iterating over rows: %v", err)
		return nil, err
	}

	return urlTags, nil
}

func (r *URLTagRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM url_tags WHERE id = $1`
	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		logger.Errorf("failed to delete url_tag: %v", err)
		return err
	}

	if result.RowsAffected() == 0 {
		logger.Warnf("No URL-Tag relationship found with ID: %s", id)
		return nil
	}
	logger.Infof("URL-Tag relationship deleted with ID: %s", id)
	return nil
}

func (r *URLTagRepository) DeleteByURLAndTag(ctx context.Context, urlID, tagID uuid.UUID) error {
	query := `DELETE FROM url_tags WHERE url_id = $1 AND tag_id = $2`
	result, err := r.db.Exec(ctx, query, urlID, tagID)
	if err != nil {
		logger.Errorf("failed to delete url_tag by url_id and tag_id: %v", err)
		return err
	}

	if result.RowsAffected() == 0 {
		logger.Warnf("No URL-Tag relationship found with URL ID: %s and Tag ID: %s", urlID, tagID)
		return nil
	}
	logger.Infof("URL-Tag relationship deleted with URL ID: %s and Tag ID: %s", urlID, tagID)
	return nil
}
