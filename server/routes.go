package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"canonical/assessment/store"
)

type BookRoutes struct {
	jsonStore store.BookJsonStore
}

func NewBookRoutes(
	group *gin.RouterGroup,
	jsonStore store.BookJsonStore) *BookRoutes {

	r := &BookRoutes{jsonStore: jsonStore}
	r.setupRoutes(group)
	return r
}

func (r *BookRoutes) setupRoutes(group *gin.RouterGroup) {
	entityGroup := group.Group("/books")
	{
		entityGroup.GET("/", r.getBooks)
		entityGroup.POST("/", r.createBook)
	}
}

// Shared
type bookResponse struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

func (r *BookRoutes) getBooks(c *gin.Context) {

	books, err := r.jsonStore.GetAll()
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

	id, err := r.jsonStore.GetNextId()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	book := store.Book{Id: id, Title: request.Title}

	if err := r.jsonStore.Set(id, book); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, bookResponse(book))
}
