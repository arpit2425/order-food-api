package dao

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

type FileDAO[T any] struct {
	FilePath string
	mu       sync.RWMutex
}

func (dao *FileDAO[T]) ReadAll() ([]T, error) {
	dao.mu.RLock()
	defer dao.mu.RUnlock()

	file, err := os.Open(dao.FilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []T{}, nil
		}
		return nil, fmt.Errorf("failed to open file %s: %w", dao.FilePath, err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", dao.FilePath, err)
	}
	if len(data) == 0 {
		return []T{}, nil
	}

	var out []T
	if err := json.Unmarshal(data, &out); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	return out, nil
}

func (dao *FileDAO[T]) WriteAll(items []T) error {
	dao.mu.Lock()
	defer dao.mu.Unlock()

	data, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}
	return ioutil.WriteFile(dao.FilePath, data, 0644)
}
