package game

import (
	"errors"
	"github.com/satori/go.uuid"
)

const startingMatches = 13
const minTakes = 1
const maxTakes = 3

// State represents the current state of a game
type State struct {
	ID          uuid.UUID
	MatchesLeft int
	Winner      string
}

// New creates a new game with no winner and an initial amount of matches
func New() State {
	return State{
		uuid.Must(uuid.NewV4()),
		startingMatches,
		"",
	}
}

// Play allows to a number of matches and let the computer play afterwards
func Play(game State, takeMatches int) (State, error) {
	if takeMatches < minTakes || takeMatches > maxTakes || takeMatches > game.MatchesLeft {
		return State{}, errors.New("invalid move")
	}

	game.MatchesLeft -= takeMatches
	if game.MatchesLeft == 0 {
		game.Winner = "Computer"
		return game, nil
	}

	game.MatchesLeft -= computerMove(game)
	if game.MatchesLeft == 0 {
		game.Winner = "User"
	}

	return game, nil
}

func computerMove(game State) int {
	if game.MatchesLeft == 1 {
		return 1
	}
	i := (game.MatchesLeft % 4) - 1
	if i <= 0 {
		return 3
	}
	return i
}
