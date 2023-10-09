package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/vnniciusg/srv_book_management/internal/app/handlers"
	"github.com/vnniciusg/srv_book_management/internal/app/models"
	"github.com/vnniciusg/srv_book_management/internal/app/usecases"
	"github.com/vnniciusg/srv_book_management/internal/infra"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := gorm.Open("postgres", "host=localhost user=postgres password=postgres dbname=library port=5434 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	db.AutoMigrate(&models.Book{}, &models.Author{})

	bookRepository := infra.NewBookRepositoryImpl(db)
	authorRepository := infra.NewAuthorRepositoryImpl(db)

	bookUseCase := &usecases.BookUseCase{BookRepository: bookRepository}
	authorUseCase := &usecases.AuthorUseCase{AuthorRepository: authorRepository}

	bookHandler := &handlers.BookHandler{BookUseCase: *bookUseCase}
	authorHandler := &handlers.AuthorHandler{AuthorUseCase: *authorUseCase}

	r := gin.Default()

	bookRoutes := r.Group("/books")
	{
		bookRoutes.POST("/", bookHandler.CreateBook)
		bookRoutes.GET("/", bookHandler.ListBook)
	}

	authorRoutes := r.Group("/authors")
	{
		authorRoutes.POST("/", authorHandler.CreateAuthor)
		authorRoutes.GET("/", authorHandler.ListAuthors)
	}

	port := 8080
	fmt.Printf("Server Listening on port %d\n", port)
	r.Run(fmt.Sprintf(":%d", port))

}
