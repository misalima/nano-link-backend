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
	query := "INSERT INTO users (id, username, email, password_hash, created_at) VALUES ($1, $2, $3, $4, $5)"
	_, err := r.db.Exec(ctx, query, user.ID, user.Username, user.Email, user.PasswordHash, user.CreatedAt)
	return err
}

func (r *UserRepository) FetchByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	query := "SELECT id, username, email, password_hash, created_at FROM users WHERE id = $1"
	row := r.db.QueryRow(ctx, query, id)

	var user domain.User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FetchByUsername(ctx context.Context, username string) (*domain.User, error) {
	query := "SELECT id, username, email, password_hash, created_at FROM users WHERE username = $1"
	row := r.db.QueryRow(ctx, query, username)

	var user domain.User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FetchByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := "SELECT id, username, email, password_hash, created_at FROM users WHERE email = $1"
	row := r.db.QueryRow(ctx, query, email)

	var user domain.User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Update(ctx context.Context, user *domain.User) error {
	query := "UPDATE users SET username = $2, email = $3, password_hash = $4 WHERE id = $1"
	_, err := r.db.Exec(ctx, query, user.ID, user.Username, user.Email, user.PasswordHash)
	return err
}

func (r *UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := r.db.Exec(ctx, query, id)
	return err
}
