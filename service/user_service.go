package service

import (
	"context"

	"github.com/rema424/twitter-reproduction-backend/model"
	"github.com/rema424/twitter-reproduction-backend/repository"
)

// UserService ...
type UserService interface {
	Get(ctx context.Context, userID int) (*model.User, error)
	SelectAll(ctx context.Context) ([]*model.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

// NewUserService ...
func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository}
}

// Get ...
func (s *userService) Get(ctx context.Context, userID int) (*model.User, error) {
	return s.userRepository.Get(ctx, userID)
}

// SelectAll ...
func (s *userService) SelectAll(ctx context.Context) ([]*model.User, error) {
	return s.userRepository.SelectAll(ctx)
}
