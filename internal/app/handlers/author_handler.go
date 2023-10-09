package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vnniciusg/srv_book_management/internal/app/models"
	"github.com/vnniciusg/srv_book_management/internal/app/usecases"
)

type AuthorHandler struct {
	AuthorUseCase usecases.AuthorUseCase
}

func (h *AuthorHandler) CreateAuthor(c *gin.Context) {
	var author models.Author

	if err := c.BindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.AuthorUseCase.CreateAuthor(&author); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, author)
}

func (h *AuthorHandler) ListAuthors(c *gin.Context) {
	authors, err := h.AuthorUseCase.ListAuthors()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, authors)
}
