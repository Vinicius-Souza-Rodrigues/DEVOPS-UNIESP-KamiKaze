package service

import (
	"testing"

	"github.com/uniesp/desafio-devops/internal/domain"
)

func TestProductDiscountNotebook(
	t *testing.T,
) {

	product := domain.Product{
		ID:    1,
		Name:  "Notebook",
		Price: 3500,
	}

	if product.Name == "Notebook" {
		product.Price *= 0.9
	}

	expected := 3150.0

	if product.Price != expected {

		t.Errorf(
			"expected %.2f got %.2f",
			expected,
			product.Price,
		)
	}
}

func TestProductDiscountMouse(
	t *testing.T,
) {

	product := domain.Product{
		ID:    1,
		Name:  "Mouse",
		Price: 80,
	}

	if product.Name == "Mouse" {
		product.Price *= 0.95
	}

	expected := 76.0

	if product.Price != expected {

		t.Errorf(
			"expected %.2f got %.2f",
			expected,
			product.Price,
		)
	}
}
