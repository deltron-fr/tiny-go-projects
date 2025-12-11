package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

// loadBookworms reads the file and returns the list of bookworms
func loadBookworms(filepath string) ([]Bookworm, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var bookworms []Bookworm

	if err := json.NewDecoder(file).Decode(&bookworms); err != nil {
		return nil, err
	}
	return bookworms, nil

}

// A Bookworm contains the list of books on a bookwarm's shelf.
type Bookworm struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

// Book describes a book on a bookwarm's shelf.
type Book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}


func findCommonBooks(bookworms []Bookworm) []Book {
	booksOnShelves := booksCount(bookworms)

	var commonBooks []Book
	for book, count := range booksOnShelves {
		if count > 1  {
			commonBooks = append(commonBooks, book)
		}
	}
	return sortBooks(commonBooks)
}

func booksCount(bookworms []Bookworm) map[Book]uint {
	count := make(map[Book]uint)

	for _, bookworm := range bookworms {
		for _, book := range bookworm.Books {
			count[book]++
		}
	}
	return count
}

func sortBooks(books []Book) []Book {
	sort.Slice(books, func(i int, j int) bool {
		if books[i].Author != books[j].Author {
			return books[i].Author < books[j].Author
		}
		return books[i].Title < books[j].Title
	})
	return books
}

func displayBooks(books []Book) {
	for _, book := range books {
		fmt.Println("-", book.Title, "by", book.Author)
	}
	fmt.Println()
}