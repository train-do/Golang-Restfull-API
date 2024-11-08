package service

import (
	"github.com/train-do/Golang-Restfull-API/model"
	"github.com/train-do/Golang-Restfull-API/repository"
)

type BookService struct {
	repo *repository.BookRepository
}

func NewBookService(repo *repository.BookRepository) *BookService {
	return &BookService{repo}
}

func (s *BookService) GetAllBook() ([]model.Book, error) {
	return s.repo.GetAll()
}

// func (s *BookService) GetProductByID(id int) (*model.Book, error) {
// 	return s.repo.GetByID(id)
// }

func (s *BookService) CreateBook(book *model.Book) error {
	return s.repo.Create(book)
}

// func (s *BookService) UpdateProduct(product *model.Book) error {
// 	return s.repo.Update(product)
// }

// func (s *BookService) DeleteProduct(id int) error {
// 	return s.repo.Delete(id)
// }
