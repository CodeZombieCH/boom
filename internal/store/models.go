package store

import "time"

type Book struct {
	ID              uint
	Title           string
	Author          *string
	PublicationDate *time.Time
	Edition         *string
	Description     *string
	Genre           *string
}
