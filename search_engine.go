package docdb

func NewSearchEngine(filePath string) *SearchEngine {
	docs := &DocumentStore{documents: make(map[int]string)}
	storage := NewPersistentStorage(filePath)
	idx := NewIndexer(storage, docs)
	return &SearchEngine{
		indexer:         idx,
		queryEngine:     NewQueryEngine(idx.invertedIndex),
		documentStore:   docs,
		persistentStore: storage,
	}
}

func (se *SearchEngine) IndexDocument(doc Document) {
	se.indexer.AddDocument(doc, se.persistentStore, se.documentStore)
}

func (se *SearchEngine) Search(query string) []int {
	return se.queryEngine.Search(query)
}
