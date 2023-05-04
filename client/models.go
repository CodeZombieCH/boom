package client

import "time"

type Book struct {
	ID              int        `json:"id"`
	Title           string     `json:"title"`
	Author          *string    `json:"author"`
	PublicationDate *time.Time `json:"publicationDate"`
	Edition         *string    `json:"edition"`
	Description     *string    `json:"description"`
	Genre           *string    `json:"genre"`
}
