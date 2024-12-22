package main

import (
	"fmt"

	docdb "github.com/satyajitnayk/doc-db"
)

func main() {
	engine := docdb.NewSearchEngine("data/index.json")

	engine.IndexDocument(docdb.Document{ID: 1, Content: "Go is a great programming language"})
	engine.IndexDocument(docdb.Document{ID: 2, Content: "Go supports concurrency and scalability"})

	results := engine.Search("Go AND scalability")
	fmt.Println("Search Results:", results) // Expected: [2]
}
