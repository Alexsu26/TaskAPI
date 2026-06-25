package service

import (
	"errors"
	"strings"

	"taskapi/internal/auth"
	"taskapi/internal/model"
	"taskapi/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo         *repository.UserRepo
	tokenManager *auth.TokenManager
}

func NewUserService(repo *repository.UserRepo, tokenManager *auth.TokenManager) *UserService {
	return &UserService{repo: repo, tokenManager: tokenManager}
}

var (
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrParaMiss           = errors.New("missing parameter")
	ErrPasswordInvalid    = errors.New("password invalid")
	ErrInvalidCredentials = errors.New("invalid email or wrong password")
	ErrTokenInvalid       = errors.New("auth failed, token generate failed")
)

func (s *UserService) Create(name, email, password string) (*model.User, error) {
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)
	password = strings.TrimSpace(password)
	if name == "" || email == "" || password == "" {
		return nil, ErrParaMiss
	}
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, ErrPasswordInvalid
	}
	user := &model.User{
		Email:        email,
		Name:         name,
		PasswordHash: string(hashBytes),
	}
	err = s.repo.Create(user)
	if err != nil {
		if errors.Is(err, repository.ErrUserEmailExists) {
			return nil, ErrEmailAlreadyExists
		}
		return nil, err
	}
	return user, nil
}

func (s *UserService) Login(email, password string) (*model.User, string, error) {
	email = strings.TrimSpace(email)
	password = strings.TrimSpace(password)
	if email == "" || password == "" {
		return nil, "", ErrParaMiss
	}
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return nil, "", ErrInvalidCredentials
		}
		return nil, "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, "", ErrInvalidCredentials
	}
	user.PasswordHash = ""
	token, err := s.tokenManager.GenerateToken(user.ID)
	if err != nil {
		if errors.Is(err, auth.ErrTokenInvalid) {
			return nil, "", ErrTokenInvalid
		}
		return nil, "", err
	}
	return user, token, nil
}
