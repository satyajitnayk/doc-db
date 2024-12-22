// Handles the process of tokenizing documents and updating the inverted index.
package docdb

import "strings"

func NewIndexer(storage *PersistentStorage, docs *DocumentStore) *Indexer {
	idx := &Indexer{
		invertedIndex: &InvertedIndex{
			index: make(map[string]map[int]bool),
		},
	}
	// Load data on startup
	storage.LoadIndex(idx.invertedIndex, docs)
	return idx
}

func tokenize(content string) []string {
	content = strings.ToLower(content)
	return strings.Fields(content)
}

func (idx *Indexer) AddDocument(doc Document, storage *PersistentStorage, docs *DocumentStore) {
	docs.mu.Lock()
	docs.documents[doc.ID] = doc.Content
	docs.mu.Unlock()

	tokens := tokenize(doc.Content)
	idx.invertedIndex.mu.Lock()
	for _, token := range tokens {
		if _, exists := idx.invertedIndex.index[token]; !exists {
			idx.invertedIndex.index[token] = make(map[int]bool)
		}
		idx.invertedIndex.index[token][doc.ID] = true
	}
	idx.invertedIndex.mu.Unlock()

	// Save the updated state to disk
	storage.SaveIndex(idx.invertedIndex, docs)
}
