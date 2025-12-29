package getstatus

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/deltron-fr/tiny-go-projects/httpgordle/internal/api"
	"github.com/go-chi/chi/v5"
)

func Handle(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, api.GameID)
	if id == "" {
		http.Error(writer, "missing the id of the game", http.StatusNotFound)
		return
	}
	log.Printf("retrieve status of game with id: %v", id)

	apiGame := api.GameResponse{
		ID: id,
	}

	err := json.NewEncoder(writer).Encode(apiGame)
	if err != nil {
		http.Error(writer, "could not encode data to json", http.StatusInternalServerError)
		return
	}
}