package searchstrategies

// SearchStrategy is an interface for different
// search types (e.g., keyword, boolean, regex).
type SearchStrategy interface {
	Search(query string) []int
}
