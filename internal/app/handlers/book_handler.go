package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vnniciusg/srv_book_management/internal/app/models"
	"github.com/vnniciusg/srv_book_management/internal/app/usecases"
)

type BookHandler struct {
	BookUseCase usecases.BookUseCase
}

func (h *BookHandler) CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.BookUseCase.CreateBook(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, book)

}

func (h *BookHandler) ListBook(c *gin.Context) {
	books, err := h.BookUseCase.ListBooks()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)

}
