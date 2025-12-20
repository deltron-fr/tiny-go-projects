package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/deltron-fr/tiny-go-projects/gordle/gordle"
)

const maxAttempts = 5

func main() {
	corpus, err := gordle.ReadCorpus("corpus/english.txt")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to read corpus: %s", err)
		return
	}

	g, err := gordle.New(bufio.NewReader(os.Stdin), corpus, maxAttempts)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to start game: %s", err)
		return
	}
	g.Play()
}