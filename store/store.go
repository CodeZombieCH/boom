package store

import "errors"

type BookStore interface {
	Set(entity *Book) (*Book, error)
	Get(id uint) (*Book, error)
	Remove(id uint) error
	GetAll() ([]Book, error)
}

type Book struct {
	ID    uint
	Title string
}

var (
	ErrNotFound = errors.New("not found")
)
