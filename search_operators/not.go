package searchoperators

// applies the NOT operation between
// a set and the entire document set.
type NOTOperator struct {
	allDocs map[int]bool
}

// performs the NOT operation on a single set of document IDs.
func (n *NOTOperator) Apply(left, _ map[int]bool) map[int]bool {
	result := make(map[int]bool)
	for docID := range n.allDocs {
		if _, exists := left[docID]; !exists {
			result[docID] = true
		}
	}
	return result
}
