package postgres

import (
	"context"
	"github.com/google/uuid"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/misalima/nano-link-backend/src/core/domain"
	"github.com/misalima/nano-link-backend/src/core/ports"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) ports.UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Save(ctx context.Context, user *domain.User) error {
	panic("unimplemented")
}

func (r *UserRepository) FetchByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	panic("unimplemented")
}

func (r *UserRepository) FetchByUsername(ctx context.Context, username string) (*domain.User, error) {
	panic("unimplemented")
}

func (r *UserRepository) FetchByEmail(ctx context.Context, email string) (*domain.User, error) {
	panic("unimplemented")
}

func (r *UserRepository) Update(ctx context.Context, user *domain.User) error {
	panic("unimplemented")
}

func (r *UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}
