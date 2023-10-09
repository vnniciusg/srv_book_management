package usecases

import (
	"github.com/vnniciusg/srv_book_management/internal/app/models"
	"github.com/vnniciusg/srv_book_management/internal/domain"
)

type AuthorUseCase struct {
	AuthorRepository domain.AuthorRepository
}

func NewAuthorUseCase(AuthorRepo domain.AuthorRepository) *AuthorUseCase {
	return &AuthorUseCase{
		AuthorRepository: AuthorRepo,
	}
}

func (uc *AuthorUseCase) CreateAuthor(author *models.Author) error {
	return uc.AuthorRepository.CreateAuthor(author)
}

func (uc *AuthorUseCase) ListAuthors() ([]*models.Author, error) {
	return uc.AuthorRepository.ListAuthors()
}
