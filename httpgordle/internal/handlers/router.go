package handlers

import (
	"github.com/go-chi/chi/v5"

	"github.com/deltron-fr/tiny-go-projects/httpgordle/internal/api"
	"github.com/deltron-fr/tiny-go-projects/httpgordle/internal/handlers/getstatus"
	"github.com/deltron-fr/tiny-go-projects/httpgordle/internal/handlers/guess"
	"github.com/deltron-fr/tiny-go-projects/httpgordle/internal/handlers/newgame"
)

func NewRouter() chi.Router {
	r := chi.NewRouter()

	r.Post(api.NewGameRoute, newgame.Handle)
	r.Get(api.GetStatusRoute, getstatus.Handle)
	r.Put(api.GuessRoute, guess.Handle)
	return r
}
