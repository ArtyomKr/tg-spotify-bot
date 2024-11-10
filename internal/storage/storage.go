package storage

import (
	"encoding/json"
	"errors"
	"io/fs"
	"os"
	"path/filepath"
)

func NewStorage(path string) (*FileStorage, error) {
	storage := &FileStorage{
		data:     make(map[string]UserData),
		filepath: path,
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 755); err != nil {
		return nil, err
	}

	if _, err := os.Stat(path); !errors.Is(err, fs.ErrNotExist) {
		file, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}

		if len(file) > 0 {
			if err := json.Unmarshal(file, &storage.data); err != nil {
				return nil, err
			}
		}
	}

	return storage, nil
}

func (s *FileStorage) Get(userID string) (UserData, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	data, ok := s.data[userID]
	return data, ok
}

func (s *FileStorage) Set(userID string, data UserData) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[userID] = data
	return s.saveToFile()
}

func (s *FileStorage) Delete(userID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, userID)
	return s.saveToFile()
}

func (s *FileStorage) saveToFile() error {
	data, err := json.Marshal(s.data)
	if err != nil {
		return err
	}
	return os.WriteFile(s.filepath, data, 0644)
}
