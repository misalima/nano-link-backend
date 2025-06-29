package postgres

import (
	"context"

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

func (r *UserRepository) Create(ctx context.Context, user *domain.User) error {
	panic("unimplemented")
}

func (r *UserRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	panic("unimplemented")
}

func (r *UserRepository) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	panic("unimplemented")
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	panic("unimplemented")
}

func (r *UserRepository) Update(ctx context.Context, user *domain.User) error {
	panic("unimplemented")
}

func (r *UserRepository) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}