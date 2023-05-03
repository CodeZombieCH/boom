package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"canonical/assessment/store"
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
	ID    uint   `json:"id"`
	Title string `json:"title"`
}

func (r *BookRoutes) getBooks(c *gin.Context) {

	books, err := r.store.GetAll()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, books)
}

type createAliasRequest struct {
	Title string `json:"title"`
}

func (r *BookRoutes) createBook(c *gin.Context) {

	var request createAliasRequest
	if err := c.BindJSON(&request); err != nil {
		return
	}

	book := &store.Book{ID: 0, Title: request.Title}

	book, err := r.store.Set(book)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, bookResponse(*book))
}
