package domain

import "github.com/vnniciusg/srv_book_management/internal/app/models"

type BookRepository interface {
	CreateBook(book *models.Book) error
	ListBooks() ([]*models.Book, error)
}
