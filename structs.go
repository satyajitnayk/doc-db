package docdb

import "sync"

type Document struct {
	ID      int
	Content string
}

type InvertedIndex struct {
	index map[string]map[int]bool // keyword -> document IDs
	mu    sync.Mutex
}

type Indexer struct {
	invertedIndex *InvertedIndex
}

type QueryEngine struct {
	invertedIndex *InvertedIndex
}

type SearchEngine struct {
	indexer         *Indexer
	queryEngine     *QueryEngine
	documentStore   *DocumentStore
	persistentStore *PersistentStorage
}

type PersistentStorage struct {
	filePath string
	mu       sync.Mutex
}

type DocumentStore struct {
	documents map[int]string // ID -> content
	mu        sync.Mutex
}
