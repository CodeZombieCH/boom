package store

import (
	"errors"
	"time"
)

type BookStore interface {
	Set(entity *Book) (*Book, error)
	Get(id uint) (*Book, error)
	Remove(id uint) error
	GetAll() ([]Book, error)
}

type Book struct {
	ID              uint
	Title           string
	Author          string
	PublicationDate time.Time
	Edition         string
	Description     string
	Genre           string
}

var (
	ErrNotFound = errors.New("store: record not found")
)
