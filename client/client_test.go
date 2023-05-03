package client

import (
	"fmt"
	"testing"
)

func TestCreateBook(t *testing.T) {
	client := NewApiClient("http://localhost:8080/api")
	book, err := client.CreateBook(Book{Title: "wuuzaa"})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(book)
}

func TestListBooks(t *testing.T) {
	client := NewApiClient("http://localhost:8080/api")
	books, err := client.ListBooks()
	if err != nil {
		t.Fatal(err)
	}

	for _, b := range books {
		fmt.Println(b)
	}
}
