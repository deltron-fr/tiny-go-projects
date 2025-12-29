package api

const (
	NewGameRoute = "/games"
	GameID = "id"
	GetStatusRoute = "/games/{" + GameID + "}"
)

// GameResponse contains the information about a game.
type GameResponse struct {
	ID string `json:"ID"`
	AttemptsLeft byte `json:"attempts_left"`
	Guesses []Guess `json:"guesses"`
	WordLength byte `json:"word_length"`
	Solution string `json:"solution,omitempty"`
	Status string `json:"status"`
}

type Guess struct {
	Word string `json:"word"`
	Feedback string `json:"feedback"`
}