package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"canonical/assessment/server"
	"canonical/assessment/store"
)

func main() {
	//store := store.BookJsonStore{StoreFilePath: "data/books.json"}

	db, err := gorm.Open(sqlite.Open("data/boom.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	store, err := store.NewBookDBStore(db)
	if err != nil {
		panic(fmt.Sprintf("failed to create store: %v", err))
	}

	engine := gin.Default()

	apiGroup := engine.Group("/api")
	server.NewBookRoutes(apiGroup, store)

	engine.Run()
}
