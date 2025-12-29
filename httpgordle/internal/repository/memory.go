package repository

import (
	"fmt"

	"github.com/deltron-fr/tiny-go-projects/httpgordle/internal/session"
)

type GameRepository struct {
	storage map[session.GameID]session.Game
}

func New() *GameRepository {
	return &GameRepository{
		storage: make(map[session.GameID]session.Game),
	}
}

// Add inserts for the first time a game in memory.
func (gr *GameRepository) Add(game session.Game) error {
	_, ok := gr.storage[game.ID]
	if ok {
		return fmt.Errorf("gameID %s already exists", game.ID)
	}
	gr.storage[game.ID] = game
	return nil
}
