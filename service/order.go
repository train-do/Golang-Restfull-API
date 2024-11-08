package service

import (
	"github.com/train-do/Golang-Restfull-API/model"
	"github.com/train-do/Golang-Restfull-API/repository"
)

type OrderService struct {
	repo *repository.OrderRepository
}

func NewOrderService(repo *repository.OrderRepository) *OrderService {
	return &OrderService{repo}
}

func (s *OrderService) GetAllOrder() ([]model.Order, error) {
	return s.repo.GetAll()
}

// func (s *OrderService) GetProductByID(id int) (*model.Order, error) {
// 	return s.repo.GetByID(id)
// }

func (s *OrderService) CreateOrder(order *model.BodyOrder) (model.Response, error) {
	return s.repo.Create(order)
}

// func (s *OrderService) UpdateProduct(product *model.Order) error {
// 	return s.repo.Update(product)
// }

// func (s *OrderService) DeleteProduct(id int) error {
// 	return s.repo.Delete(id)
// }
