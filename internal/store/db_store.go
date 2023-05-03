package store

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type BookDBStore struct {
	db *gorm.DB
}

func NewBookDBStore(db *gorm.DB) (*BookDBStore, error) {
	err := db.AutoMigrate(&book{})
	if err != nil {
		return nil, fmt.Errorf("store: migration failed: %w", err)
	}
	return &BookDBStore{db: db}, nil
}

func (s *BookDBStore) Set(entity *Book) (*Book, error) {
	model := book{
		Model:           gorm.Model{ID: entity.ID},
		Title:           entity.Title,
		Author:          entity.Author,
		PublicationDate: entity.PublicationDate,
		Edition:         entity.Edition,
		Description:     entity.Description,
		Genre:           entity.Genre,
	}

	result := s.db.Save(&model)
	if result.Error != nil {
		return nil, fmt.Errorf("store: failed to save record: %w", result.Error)
	}

	return s.Get(model.ID)
}

func (s *BookDBStore) Get(id uint) (*Book, error) {
	var model book

	if result := s.db.First(&model, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		} else {
			return nil, fmt.Errorf("store: failed to get record: %w", result.Error)
		}
	}
	return &Book{
		ID:              model.ID,
		Title:           model.Title,
		Author:          model.Author,
		PublicationDate: model.PublicationDate,
		Edition:         model.Edition,
		Description:     model.Description,
		Genre:           model.Genre,
	}, nil
}

func (s *BookDBStore) Remove(id uint) error {
	if result := s.db.Delete(&book{}, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return ErrNotFound
		} else {
			return fmt.Errorf("store: failed to delete record: %w", result.Error)
		}
	}

	return nil
}

func (s *BookDBStore) GetAll() ([]Book, error) {
	var models []book
	result := s.db.Find(&models)
	if result.Error != nil {
		return nil, fmt.Errorf("store: failed to get records: %w", result.Error)
	}

	var books []Book
	for i := 0; i < len(models); i++ {
		model := models[i]
		books = append(books, Book{
			ID:              model.ID,
			Title:           model.Title,
			Author:          model.Author,
			PublicationDate: model.PublicationDate,
			Edition:         model.Edition,
			Description:     model.Description,
			Genre:           model.Genre,
		})
	}

	return books, nil
}

// Internal model of a book used by the ORM
type book struct {
	gorm.Model
	Title           string
	Author          *string
	PublicationDate *time.Time
	Edition         *string
	Description     *string
	Genre           *string
}
