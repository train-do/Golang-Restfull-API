package repository

import (
	"database/sql"

	"github.com/train-do/Golang-Restfull-API/model"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db}
}

func (r *OrderRepository) GetAll() ([]model.Order, error) {
	rows, err := r.db.Query(`SELECT * FROM "Order"`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := []model.Order{}
	for rows.Next() {
		var order model.Order
		if err := rows.Scan(&order.Id, &order.OrderCode, &order.Name, &order.TotalAmount, &order.FinalAmount, &order.Status, &order.CreatedAt); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (r *OrderRepository) Create(order *model.BodyOrder) (model.Response, error) {
	respon := model.Response{}
	r.db.Exec("INSERT INTO orders (name, price) VALUES ($1, $2)", order.Name, order.Price)
	return respon, nil
}

// func (r *OrderRepository) GetByID(id int) (*model.Order, error) {
// 	var order model.Order
// 	err := r.db.QueryRow("SELECT id, name, price FROM orders WHERE id=$1", id).Scan(&order.ID, &order.Name, &order.Price)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &order, nil
// }
