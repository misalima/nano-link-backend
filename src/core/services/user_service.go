package services

import (
	"context"
	"github.com/google/uuid"

	"github.com/misalima/nano-link-backend/src/core/domain"
	"github.com/misalima/nano-link-backend/src/core/ports"
)

type UserService struct {
	userRepo ports.UserRepository
}

func NewUserService(userRepo ports.UserRepository) ports.UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) Register(ctx context.Context, username, email, password string) (*domain.User, error) {
	panic("unimplemented")
}

func (s *UserService) Authenticate(ctx context.Context, usernameOrEmail, password string) (*domain.User, error) {
	panic("unimplemented")
}

func (s *UserService) GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	panic("unimplemented")
}

func (s *UserService) UpdateUser(ctx context.Context, user *domain.User) error {
	panic("unimplemented")
}

func (s *UserService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}