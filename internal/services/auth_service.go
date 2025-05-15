package services

import (
	"errors"

	"draw/internal/models"
	"draw/internal/repository"
)

var (
	ErrInvalidCredentials = errors.New("invalid username or password")
	ErrUsernameTaken      = errors.New("username already taken")
	ErrEmailTaken         = errors.New("email already taken")
)

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Register(username, email, password string) (*models.User, error) {
	// Check if username is already taken
	_, err := s.userRepo.FindByUsername(username)
	if err == nil {
		return nil, ErrUsernameTaken
	} else if err != repository.ErrUserNotFound {
		return nil, err
	}

	// Check if email is already taken
	_, err = s.userRepo.FindByEmail(email)
	if err == nil {
		return nil, ErrEmailTaken
	} else if err != repository.ErrUserNotFound {
		return nil, err
	}

	// Create new user
	user := &models.User{
		Username: username,
		Email:    email,
		Password: password,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) Login(username, password string) (*models.User, error) {
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		if err == repository.ErrUserNotFound {
			return nil, ErrInvalidCredentials
		}
		return nil, err
	}

	if !user.CheckPassword(password) {
		return nil, ErrInvalidCredentials
	}

	return user, nil
}
