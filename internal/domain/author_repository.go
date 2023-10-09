package domain

import "github.com/vnniciusg/srv_book_management/internal/app/models"

type AuthorRepository interface {
	CreateAuthor(author *models.Author) error
	ListAuthors() ([]*models.Author, error)
}
