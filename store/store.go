package store

import (
	"encoding/json"
	"errors"
	"os"
)

var (
	ErrNotFound = errors.New("not found")
)

type Book struct {
	Id    int
	Title string
}

type BookJsonStore struct {
	StoreFilePath string
}

func (s BookJsonStore) GetNextId() (int, error) {
	entities, err := s.load()
	if err != nil {
		return -1, err
	}

	maxId := 0
	for _, e := range entities {
		if e.Id > maxId {
			maxId = e.Id
		}
	}

	return maxId + 1, nil
}

func (s BookJsonStore) Set(id int, entity Book) error {
	entities, err := s.load()
	if err != nil {
		return err
	}

	// Try to update existing
	var found bool
	for i := range entities {
		if entities[i].Id == id {
			entities[i] = entity
			found = true
			break
		}
	}

	if !found {
		// Create new
		entities = append(entities, entity)
	}

	if err := s.save(entities); err != nil {
		return err
	}

	return nil
}

func (s BookJsonStore) Get(id int) (Book, error) {
	entities, err := s.load()
	if err != nil {
		return Book{}, err
	}

	for i := range entities {
		if entities[i].Id == id {
			return entities[i], nil
		}
	}

	return Book{}, ErrNotFound
}

func (s BookJsonStore) Remove(id int) error {
	entities, err := s.load()
	if err != nil {
		return err
	}

	// Search for index
	indexToDelete := -1
	for i := range entities {
		if entities[i].Id == id {
			indexToDelete = i
			break
		}
	}

	if indexToDelete == -1 {
		return ErrNotFound
	}

	entities = append(entities[:indexToDelete], entities[indexToDelete+1:]...)

	if err := s.save(entities); err != nil {
		return err
	}

	return nil
}

func (s BookJsonStore) GetAll() ([]Book, error) {
	aliases, err := s.load()
	if err != nil {
		return nil, err
	}

	return aliases, nil
}

func (s BookJsonStore) load() ([]Book, error) {
	file, err := os.ReadFile(s.StoreFilePath)
	if err != nil {
		return nil, err
	}

	data := []Book{}

	if err := json.Unmarshal([]byte(file), &data); err != nil {
		return nil, err
	}

	return data, nil
}

func (s BookJsonStore) save(data []Book) error {
	bytes, err := json.MarshalIndent(data, "", "  ")

	if err != nil {
		return err
	}

	if err := os.WriteFile(s.StoreFilePath, bytes, 0644); err != nil {
		return err
	}

	return nil
}
