package main

import (
	"net/http"

	"github.com/deltron-fr/tiny-go-projects/httpgordle/internal/handlers"
)

func main() {
	err := http.ListenAndServe(":8080", handlers.NewRouter())
	if err != nil {
		panic(err)
	}
}
