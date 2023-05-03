package store

import (
	"encoding/json"
	"os"
)

type BookJsonStore struct {
	StoreFilePath string
}

func (s *BookJsonStore) getNextId() (uint, error) {
	entities, err := s.load()
	if err != nil {
		return 0, err
	}

	var maxId uint = 0
	for _, e := range entities {
		if e.ID > maxId {
			maxId = e.ID
		}
	}

	return maxId + 1, nil
}

func (s *BookJsonStore) Set(entity Book) error {
	entities, err := s.load()
	if err != nil {
		return err
	}

	if entity.ID > 0 {
		// Update
		var found bool
		for i := range entities {
			if entities[i].ID == entity.ID {
				entities[i] = entity
				found = true
				break
			}
		}

		if !found {
			return ErrNotFound
		}
	} else {
		// Create
		entity.ID, err = s.getNextId()
		if err != nil {
			return err
		}
		entities = append(entities, entity)
	}

	if err := s.save(entities); err != nil {
		return err
	}

	return nil
}

func (s *BookJsonStore) Get(id uint) (Book, error) {
	entities, err := s.load()
	if err != nil {
		return Book{}, err
	}

	for i := range entities {
		if entities[i].ID == id {
			return entities[i], nil
		}
	}

	return Book{}, ErrNotFound
}

func (s *BookJsonStore) Remove(id uint) error {
	entities, err := s.load()
	if err != nil {
		return err
	}

	// Search for index
	indexToDelete := -1
	for i := range entities {
		if entities[i].ID == id {
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

func (s *BookJsonStore) GetAll() ([]Book, error) {
	aliases, err := s.load()
	if err != nil {
		return nil, err
	}

	return aliases, nil
}

func (s *BookJsonStore) load() ([]Book, error) {
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

func (s *BookJsonStore) save(data []Book) error {
	bytes, err := json.MarshalIndent(data, "", "  ")

	if err != nil {
		return err
	}

	if err := os.WriteFile(s.StoreFilePath, bytes, 0644); err != nil {
		return err
	}

	return nil
}
