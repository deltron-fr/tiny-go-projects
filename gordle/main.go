package main

import (
	"os"

	"github.com/deltron-fr/tiny-go-projects/gordle/gordle"
)

const maxAttempts = 3

func main() {
	solution := "Hello"

	g := gordle.New(os.Stdin, solution, maxAttempts)
	g.Play()
}