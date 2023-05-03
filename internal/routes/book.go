package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"canonical/assessment/internal/store"
)

type BookRoutes struct {
	store store.BookStore
}

func NewBookRoutes(
	group *gin.RouterGroup,
	store store.BookStore) *BookRoutes {

	r := &BookRoutes{store: store}
	r.setupRoutes(group)
	return r
}

func (r *BookRoutes) setupRoutes(group *gin.RouterGroup) {
	entityGroup := group.Group("/books")
	{
		entityGroup.GET("", r.getBooks)
		entityGroup.POST("", r.createBook)
	}
}

// Shared
type bookResponse struct {
	ID              uint       `json:"id"`
	Title           string     `json:"title"`
	Author          *string    `json:"author"`
	PublicationDate *time.Time `json:"publicationDate"`
	Edition         *string    `json:"edition"`
	Description     *string    `json:"description"`
	Genre           *string    `json:"genre"`
}

func (r *BookRoutes) getBooks(c *gin.Context) {
	books, err := r.store.GetAll()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, books)
}

type createBookRequest struct {
	Title           string     `json:"title"`
	Author          *string    `json:"author"`
	PublicationDate *time.Time `json:"publicationDate"`
	Edition         *string    `json:"edition"`
	Description     *string    `json:"description"`
	Genre           *string    `json:"genre"`
}

func (r *BookRoutes) createBook(c *gin.Context) {
	var request createBookRequest
	if err := c.BindJSON(&request); err != nil {
		return
	}

	book := &store.Book{
		ID:              0,
		Title:           request.Title,
		Author:          request.Author,
		PublicationDate: request.PublicationDate,
		Edition:         request.Edition,
		Description:     request.Description,
		Genre:           request.Genre,
	}

	book, err := r.store.Set(book)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, bookResponse(*book))
}
