package repository

import (
	"database/sql"
	"fmt"

	"github.com/train-do/Golang-Restfull-API/model"
)

type BookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{db}
}

func (r *BookRepository) GetAll() ([]model.Book, error) {
	rows, err := r.db.Query(`SELECT * FROM "Book"`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := []model.Book{}
	for rows.Next() {
		var book model.Book
		if err := rows.Scan(&book.Id, &book.BookCode, &book.Title, &book.Type, &book.Author, &book.Price, &book.Discount); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

// func (r *BookRepository) GetByID(id int) (*model.Book, error) {
// 	var book model.Book
// 	err := r.db.QueryRow("SELECT id, name, price FROM books WHERE id=$1", id).Scan(&book.ID, &book.Name, &book.Price)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &book, nil
// }

func (r *BookRepository) Create(book *model.Book) error {
	var book_code string
	err := r.db.QueryRow(`select lpad(nextval('"Book_id_seq"')::text, 3, '0') as curval;`).Scan(&book_code)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return err
	}
	query := `INSERT INTO "Book" (book_code, title, type, author, price, discount) VALUES ($1, $2, $3, $4, $5, $6) returning id, book_code`
	err = r.db.QueryRow(query, book_code, book.Title, book.Type, book.Author, book.Price, book.Discount).Scan(&book.Id, &book.BookCode)
	// fmt.Printf("%+v------\n", book)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return err
	}
	return err
}

// func (r *BookRepository) Update(book *model.Book) error {
// 	_, err := r.db.Exec("UPDATE books SET name=$1, price=$2 WHERE id=$3", book.Name, book.Price, book.ID)
// 	return err
// }

// func (r *BookRepository) Delete(id int) error {
// 	_, err := r.db.Exec("DELETE FROM books WHERE id=$1", id)
// 	return err
// }
