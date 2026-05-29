package service

import (
	"github.com/uniesp/desafio-devops/internal/domain"
	"github.com/uniesp/desafio-devops/internal/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(
	repo *repository.ProductRepository,
) *ProductService {

	return &ProductService{
		repo: repo,
	}
}

func (s *ProductService) FindAll() (
	[]domain.Product,
	error,
) {

	products, err := s.repo.FindAll()

	if err != nil {
		return nil, err
	}

	for i := range products {

		if products[i].Name == "Notebook" {
			products[i].Price *= 0.9
		}

		if products[i].Name == "Mouse" {
			products[i].Price *= 0.95
		}
	}

	return products, nil
}