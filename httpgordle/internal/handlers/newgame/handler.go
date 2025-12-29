package newgame

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/deltron-fr/tiny-go-projects/httpgordle/internal/api"
)

func Handle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	apiGame := api.GameResponse{}
	err := json.NewEncoder(w).Encode(apiGame)
	if err != nil {
		log.Printf("failed to write response: %s", err)
	}
}