package store

import (
	"errors"

	"gorm.io/gorm"
)

type BookDBStore struct {
	db *gorm.DB
}

func NewBookDBStore(db *gorm.DB) (*BookDBStore, error) {
	err := db.AutoMigrate(&bookModel{})
	if err != nil {
		return nil, err
	}
	return &BookDBStore{db: db}, nil
}

func (s *BookDBStore) Set(entity *Book) (*Book, error) {
	model := bookModel{Title: entity.Title}

	result := s.db.Save(&model)
	if result.Error != nil {
		return nil, result.Error
	}

	return s.Get(model.ID)
}

func (s *BookDBStore) Get(id uint) (*Book, error) {
	var model bookModel

	if result := s.db.First(&model, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		} else {
			return nil, result.Error
		}
	}
	return &Book{ID: model.ID, Title: model.Title}, nil
}

func (s *BookDBStore) Remove(id uint) error {
	if result := s.db.Delete(&bookModel{}, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return ErrNotFound
		} else {
			return result.Error
		}
	}

	return nil
}

func (s *BookDBStore) GetAll() ([]Book, error) {
	var models []bookModel
	result := s.db.Find(&models)
	if result.Error != nil {
		return nil, result.Error
	}

	var books []Book
	for i := 0; i < len(models); i++ {
		model := models[i]
		books = append(books, Book{ID: model.ID, Title: model.Title})
	}

	return books, nil
}

func (s *BookDBStore) GetNextId() (uint, error) {
	return 0, nil
}

type bookModel struct {
	gorm.Model
	Title string
}
