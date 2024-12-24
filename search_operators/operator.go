package searchoperators

type SearchOpeartor interface {
	Apply(left, right map[int]bool) map[int]bool
}
