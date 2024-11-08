package repository

import (
	"database/sql"

	"github.com/train-do/Golang-Restfull-API/model"
)

type ReviewRepository struct {
	db *sql.DB
}

func NewReviewRepository(db *sql.DB) *ReviewRepository {
	return &ReviewRepository{db}
}

func (r *ReviewRepository) Create(product *model.Review) error {
	// _, err := r.db.Exec("INSERT INTO products (name, price) VALUES ($1, $2)", product.Name, product.Price)
	return nil
	// return err
}
