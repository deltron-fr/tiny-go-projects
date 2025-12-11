package main

import (
	"fmt"
	"os"
	"flag"
)

func main() {
	var filepath string

	flag.StringVar(&filepath, "filepath", "testdata/bookworms.json", "The required json file path")
	flag.Parse()

	bookworms, err := loadBookworms(filepath)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to load bookworms: %v", err)
		os.Exit(1)
	}

	commonBooks := findCommonBooks(bookworms)
	fmt.Println("Here are the books in common:")
	displayBooks(commonBooks)

	readers := getAllReaders(bookworms)
	targetReader := getReader("Fadi", bookworms)

	recommendations := recommend(readers, targetReader, 5)
	displayRecommendations(recommendations)
}
