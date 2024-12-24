package searchoperators

type ANDOperator struct{}

func (a *ANDOperator) Apply(left, right map[int]bool) map[int]bool {
	result := make(map[int]bool)
	// intersection
	for docID := range left {
		if _, exists := right[docID]; exists {
			result[docID] = true
		}
	}
	return result
}
