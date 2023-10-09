package infra

import (
	"github.com/jinzhu/gorm"
	"github.com/vnniciusg/srv_book_management/internal/app/models"
)

type AuthorRepositoryImpl struct {
	DB *gorm.DB
}

func NewAuthorRepositoryImpl(db *gorm.DB) *AuthorRepositoryImpl {
	return &AuthorRepositoryImpl{
		DB: db,
	}
}

func (r *AuthorRepositoryImpl) CreateAuthor(author *models.Author) error {
	return r.DB.Create(author).Error
}

func (r *AuthorRepositoryImpl) ListAuthors() ([]*models.Author, error) {
	var authors []*models.Author
	if err := r.DB.Find(&authors).Error; err != nil {
		return nil, err
	}

	return authors, nil
}
