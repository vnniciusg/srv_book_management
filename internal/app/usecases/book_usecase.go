package usecases

import (
	"github.com/vnniciusg/srv_book_management/internal/app/models"
	"github.com/vnniciusg/srv_book_management/internal/domain"
)

type BookUseCase struct {
	BookRepository domain.BookRepository
}

func NewBookUseCase(BookRepo domain.BookRepository) *BookUseCase {
	return &BookUseCase{
		BookRepository: BookRepo,
	}
}

func (uc *BookUseCase) CreateBook(book *models.Book) error {
	return uc.BookRepository.CreateBook(book)
}

func (uc *BookUseCase) ListBooks() ([]*models.Book, error) {
	return uc.BookRepository.ListBooks()
}
