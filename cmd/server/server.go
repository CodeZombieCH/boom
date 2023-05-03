package main

import (
	"github.com/gin-gonic/gin"

	"canonical/assessment/server"
	"canonical/assessment/store"
)

func main() {
	jsonStore := store.BookJsonStore{StoreFilePath: "data/books.json"}

	engine := gin.Default()

	apiGroup := engine.Group("/api")
	server.NewBookRoutes(apiGroup, jsonStore)

	engine.Run()
}
