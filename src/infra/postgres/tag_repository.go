package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/misalima/nano-link-backend/src/infra/logger"

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
	query := `INSERT INTO tags (id, name, created_at) VALUES ($1, $2, $3)`

	_, err := r.db.Exec(ctx, query, tag.ID, tag.Name, tag.CreatedAt)
	if err != nil {
		logger.Errorf("failed to insert tag: %v", err)
		return err
	}

	logger.Infof("Tag saved with ID: %s", tag.ID)

	return nil
}

func (r *TagRepository) FetchByID(ctx context.Context, id uuid.UUID) (*domain.Tag, error) {
	query := `SELECT id, name, created_at FROM tags WHERE id = $1`

	row := r.db.QueryRow(ctx, query, id)

	var tag domain.Tag
	err := row.Scan(&tag.ID, &tag.Name, &tag.CreatedAt)
	if err != nil {
		logger.Errorf("failed to fetch tag by id: %v", err)
		return nil, err
	}

	return &tag, nil
}

func (r *TagRepository) FetchByName(ctx context.Context, name string) (*domain.Tag, error) {
	query := `SELECT id, name, created_at FROM tags WHERE name = $1`

	row := r.db.QueryRow(ctx, query, name)

	var tag domain.Tag
	err := row.Scan(&tag.ID, &tag.Name, &tag.CreatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			logger.Infof("No tag found with name: %s", name)
			return nil, domain.ErrTagNameNotFound
		}
		logger.Errorf("failed to fetch tag by name: %v", err)
		return nil, err
	}

	return &tag, nil
}

func (r *TagRepository) FetchAll(ctx context.Context) ([]*domain.Tag, error) {
	query := `SELECT id, name, created_at FROM tags ORDER BY name`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		logger.Errorf("failed to fetch all tags: %v", err)
		return nil, err
	}
	defer rows.Close()

	var tags []*domain.Tag
	for rows.Next() {
		var tag domain.Tag
		err := rows.Scan(&tag.ID, &tag.Name, &tag.CreatedAt)
		if err != nil {
			logger.Errorf("failed to scan tag row: %v", err)
			return nil, err
		}
		tags = append(tags, &tag)
	}

	if err := rows.Err(); err != nil {
		logger.Errorf("error iterating over rows: %v", err)
		return nil, err
	}

	return tags, nil
}

func (r *TagRepository) Update(ctx context.Context, tag *domain.Tag) error {
	query := `UPDATE tags SET name = $1 WHERE id = $2`

	_, err := r.db.Exec(ctx, query, tag.Name, tag.ID)
	if err != nil {
		logger.Errorf("failed to update tag: %v", err)
		return err
	}

	logger.Infof("Tag updated with ID: %s", tag.ID)

	return nil
}

func (r *TagRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM tags WHERE id = $1`
	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		logger.Errorf("failed to delete tag: %v", err)
		return err
	}

	if result.RowsAffected() == 0 {
		logger.Warnf("No tag found with ID: %s", id)
		return nil
	}
	logger.Infof("Tag deleted with ID: %s", id)
	return nil
}
