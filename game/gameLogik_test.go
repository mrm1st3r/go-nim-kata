package game

import (
	"testing"
	"github.com/satori/go.uuid"
)

type test struct {
	MatchesLeft int
	ExpectedTake int
}

var tests = []test{
	test{1, 1},
	test{2, 1},
	test{3, 2},
	test{4, 3},
	test{6, 1},
	test{7, 2},
}

func TestComputerMove(t *testing.T) {
	for _,data := range tests {
		actual := computerMove(State{uuid.UUID{}, data.MatchesLeft, ""})
		if actual != data.ExpectedTake {
			t.Errorf("For left matches %d expected %d, got %d", data.MatchesLeft, data.ExpectedTake, actual)
		}
	}
}