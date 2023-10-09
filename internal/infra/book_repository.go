package infra

import (
	"github.com/jinzhu/gorm"
	"github.com/vnniciusg/srv_book_management/internal/app/models"
)

type BookRepositoryImpl struct {
	DB *gorm.DB
}

func NewBookRepositoryImpl(db *gorm.DB) *BookRepositoryImpl {
	return &BookRepositoryImpl{
		DB: db,
	}
}

func (r *BookRepositoryImpl) CreateBook(book *models.Book) error {
	return r.DB.Create(book).Error
}

func (r *BookRepositoryImpl) ListBooks() ([]*models.Book, error) {
	var books []*models.Book
	if err := r.DB.Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}
