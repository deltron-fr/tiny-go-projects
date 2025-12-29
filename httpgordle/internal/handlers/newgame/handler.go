package newgame

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/deltron-fr/tiny-go-projects/httpgordle/internal/api"
	"github.com/deltron-fr/tiny-go-projects/httpgordle/internal/session"
)

// Handle is the handler for the game creation endpoint.
func Handle(w http.ResponseWriter, req *http.Request) {
	game := createGame()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	apiGame := api.ToGameResponse(game)
	err := json.NewEncoder(w).Encode(apiGame)
	if err != nil {
		// The header has already been set. Nothing much we can do here.
		log.Printf("failed to write response: %s", err)
	}
}

func createGame() session.Game {
	return session.Game{}
}
