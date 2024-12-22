package docdb

import (
	"encoding/json"
	"os"
)

func NewPersistentStorage(filePath string) *PersistentStorage {
	return &PersistentStorage{filePath: filePath}
}

// SaveIndex saves the index and documents to disk.
func (ps *PersistentStorage) SaveIndex(index *InvertedIndex, docs *DocumentStore) error {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	data := struct {
		Index map[string]map[int]bool `json:"index"`
		Docs  map[int]string          `json:"documents"`
	}{
		Index: index.index,
		Docs:  docs.documents,
	}

	file, err := os.Create(ps.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(data)
}

func (ps *PersistentStorage) LoadIndex(index *InvertedIndex, docs *DocumentStore) error {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	file, err := os.Open(ps.filePath)
	if os.IsNotExist(err) {
		return nil
	} else if err != nil {
		return err
	}
	defer file.Close()

	data := struct {
		Index map[string]map[int]bool `json:"index"`
		Docs  map[int]string          `json:"documents"`
	}{}

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return err
	}

	index.index = data.Index
	docs.documents = data.Docs
	return nil
}
