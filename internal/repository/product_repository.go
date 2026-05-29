package repository

import (
	"database/sql"

	"github.com/uniesp/desafio-devops/internal/domain"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(
	db *sql.DB,
) *ProductRepository {

	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) FindAll() (
	[]domain.Product,
	error,
) {

	rows, err := r.db.Query(`
		SELECT id, name, price
		FROM products
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []domain.Product

	for rows.Next() {

		var product domain.Product

		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
		)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}