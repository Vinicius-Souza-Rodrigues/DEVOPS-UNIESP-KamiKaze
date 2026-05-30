package service

import (
	"errors"

	"github.com/uniesp/desafio-devops/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(
	repo *repository.UserRepository,
) *UserService {

	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Create(
	name string,
) error {

	if name == "" {
		return errors.New("invalid name")
	}

	return s.repo.Create(name)
}