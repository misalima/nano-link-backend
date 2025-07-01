package services

import (
	"context"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := domain.NewUser(username, email, string(hashedPassword))

	err = s.userRepo.Save(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Authenticate(ctx context.Context, usernameOrEmail, password string) (*domain.User, error) {
	user, err := s.userRepo.FetchByUsername(ctx, usernameOrEmail)
	if err != nil {
		user, err = s.userRepo.FetchByEmail(ctx, usernameOrEmail)
		if err != nil {
			return nil, err
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	return s.userRepo.FetchByID(ctx, id)
}

func (s *UserService) UpdateUser(ctx context.Context, user *domain.User) error {
	return s.userRepo.Update(ctx, user)
}

func (s *UserService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return s.userRepo.Delete(ctx, id)
}