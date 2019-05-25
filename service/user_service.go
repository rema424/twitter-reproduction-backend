package service

import (
	"github.com/rema424/twitter-reproduction-backend/model"
	"github.com/rema424/twitter-reproduction-backend/repository"
)

// UserService ...
type UserService interface {
	Get(userID int) (*model.User, error)
	SelectAll() ([]*model.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

// NewUserService ...
func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository}
}

// Get ...
func (s *userService) Get(userID int) (*model.User, error) {
	return s.userRepository.Get(userID)
}

// SelectAll ...
func (s *userService) SelectAll() ([]*model.User, error) {
	return s.userRepository.SelectAll()
}
