package game

import (
	"github.com/satori/go.uuid"
	"testing"
)

type test struct {
	MatchesLeft  int
	ExpectedTake int
}

var tests = []test{
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{6, 1},
	{7, 2},
}

func TestComputerMove(t *testing.T) {
	for _, data := range tests {
		actual := computerMove(State{uuid.UUID{}, data.MatchesLeft, ""})
		if actual != data.ExpectedTake {
			t.Errorf("For left matches %d expected %d, got %d", data.MatchesLeft, data.ExpectedTake, actual)
		}
	}
}
