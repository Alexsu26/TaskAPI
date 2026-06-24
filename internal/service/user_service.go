package service

import (
	"errors"
	"strings"

	"taskapi/internal/model"
	"taskapi/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepo
}

func NewUserService(repo *repository.UserRepo) *UserService {
	return &UserService{repo: repo}
}

var (
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrParaMiss           = errors.New("missing parameter")
	ErrPasswordInvalid    = errors.New("password invalid")
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
