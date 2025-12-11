package main

import (
	"fmt"
	"math"
	"sort"
)


type Recommendation struct {
	Book Book
	Score float64
}

type Reader struct {
	Name string
	Books []Book
}

type set map[Book]struct{}

type recommendationPair struct {
	Book Book
	Similarity float64
}

func recommend(allReaders []Reader, target Reader, n int) []Recommendation {
	read := newSet(target.Books...)

	recommendations := map[Book]float64{}
	for _, reader := range allReaders {
		if reader.Name == target.Name {
			continue
		}

		var similarity float64
		for _, book := range reader.Books {
			if read.Contains(book) {
				similarity++
			}
		}

		if similarity == 0 {
			continue
		}

		score := math.Log(similarity) + 1
		for _, book := range reader.Books {
			if !read.Contains(book) {
				recommendations[book] += score
			}
		}
	}
	sortedRecommendations := getFirstNElements(sortRecommendedBooks(recommendations), n)
	var allRecommendations []Recommendation

	for book, score := range sortedRecommendations {
		rec := Recommendation{
			Book: book,
			Score: score,
		}
		allRecommendations = append(allRecommendations, rec)
	}

	return allRecommendations
}

func getAllReaders(bookworms []Bookworm) []Reader {
	var readers []Reader

	for _, bookworm := range bookworms {
		readers = append(readers, Reader(bookworm))
	}
	return readers
}

func getReader(reader string, bookworms []Bookworm) Reader {
	for _, bookworm := range bookworms {
		if bookworm.Name == reader { 
			return Reader(bookworm)
		}
	}
	return Reader{}
}

func (s set) Contains(b Book) bool {
	_, ok := s[b]
	return ok
} 

func newSet(books ...Book) set {
	m := make(map[Book]struct{})

	for _, book := range books {
		m[book] = struct{}{}
	}

	return m
}

func sortRecommendedBooks(recommendations map[Book]float64) map[Book]float64 {
	var kvPair []recommendationPair

	for k, v := range recommendations {
		kvPair = append(kvPair, recommendationPair{k, v})
	}

	sort.Slice(kvPair, func(i, j int) bool {
		if kvPair[i].Similarity != kvPair[j].Similarity {
			return kvPair[i].Similarity > kvPair[j].Similarity
		}
		return kvPair[i].Book.Title < kvPair[j].Book.Title
	})

	recommendationsSorted := make(map[Book]float64)

	for _, kv := range kvPair {
		recommendationsSorted[kv.Book] = kv.Similarity
	}

	return recommendationsSorted
}

func getFirstNElements(recommendations map[Book]float64, n int) map[Book]float64 {

	count := 0
	firstNrecommendations := make(map[Book]float64)
	for k, v := range recommendations {
		if count < n {
			firstNrecommendations[k] = v
		} else {
			break
		}
		count++
	}
	return firstNrecommendations
}

func displayRecommendations(recommendations []Recommendation) {
	for _, rec := range recommendations {
		fmt.Println("Recommended Books: ")
		fmt.Println("Book Title:", rec.Book.Title, "by", rec.Book.Author, "with score:", rec.Score)
	}
}