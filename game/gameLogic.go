package game

import (
	"errors"
	"github.com/satori/go.uuid"
)

const startingMatches = 13
const minTakes = 1
const maxTakes = 3

type State struct {
	ID          uuid.UUID
	MatchesLeft int
	Winner      string
}

func New() State {
	return State{
		uuid.Must(uuid.NewV4()),
		startingMatches,
		"",
	}
}

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
