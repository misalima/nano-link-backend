package postgres

import (
	"context"

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
	panic("unimplemented")
}

func (r *URLRepository) FetchByID(ctx context.Context, id string) (*domain.URL, error) {
	panic("unimplemented")
}

func (r *URLRepository) FetchByShortID(ctx context.Context, shortID string) (*domain.URL, error) {
	panic("unimplemented")
}

func (r *URLRepository) FetchByCustomShortID(ctx context.Context, customShortID string) (*domain.URL, error) {
	panic("unimplemented")
}

func (r *URLRepository) FetchByUserID(ctx context.Context, userID string) ([]*domain.URL, error) {
	panic("unimplemented")
}

func (r *URLRepository) Update(ctx context.Context, url *domain.URL) error {
	panic("unimplemented")
}

func (r *URLRepository) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}
