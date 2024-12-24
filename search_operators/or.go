package searchoperators

type OROperator struct{}

func (o *OROperator) Apply(left, right map[int]bool) map[int]bool {
	result := make(map[int]bool)
	// Union of both sets
	for docID := range left {
		result[docID] = true
	}
	for docID := range right {
		result[docID] = true
	}
	return result
}
