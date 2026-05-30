package repository

import "database/sql"

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(
	db *sql.DB,
) *UserRepository {

	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(
	name string,
) error {

	_, err := r.db.Exec(
		"INSERT INTO users (name) VALUES ($1)",
		name,
	)

	return err
}