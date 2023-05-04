package store

import (
	"errors"
)

type BookStore interface {
	Set(entity *Book) (*Book, error)
	Get(id uint) (*Book, error)
	Remove(id uint) error
	GetAll() ([]Book, error)
}

var (
	ErrNotFound = errors.New("store: record not found")
)
