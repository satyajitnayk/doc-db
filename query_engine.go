package docdb

func NewQueryEngine(index *InvertedIndex) *QueryEngine {
	return &QueryEngine{invertedIndex: index}
}

func (qe *QueryEngine) Search(query string) []int {
	tokens := tokenize(query)
	result := make(map[int]bool)

	for i, token := range tokens {
		if docIDs, exists := qe.invertedIndex.index[token]; exists {
			if i == 0 {
				// Initialize result with first token
				for docID := range docIDs {
					result[docID] = true
				}
			} else {
				// perform AND operation
				for docID := range result {
					if !docIDs[docID] {
						delete(result, docID)
					}
				}
			}
		} else {
			return nil // If any token doesn't exist, result is empty
		}
	}

	// convert map keys to slice
	docs := []int{}
	for docID := range result {
		docs = append(docs, docID)
	}
	return docs
}
