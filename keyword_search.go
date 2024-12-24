package docdb

import "strings"

type KeywordSearchStrategy struct {
	invertedIndex *InvertedIndex
}

func NewKeywordSearchStrategy(index *InvertedIndex) *KeywordSearchStrategy {
	return &KeywordSearchStrategy{invertedIndex: index}
}

func (k *KeywordSearchStrategy) Search(query string) []int {
	query = strings.ToLower(query)
	k.invertedIndex.mu.Lock()
	defer k.invertedIndex.mu.Unlock()

	// Return the list of document IDs for the keyword
	if docIDs, found := k.invertedIndex.index[query]; found {
		var result []int
		for docID := range docIDs {
			result = append(result, docID)
		}
		return result
	}
	return nil
}
