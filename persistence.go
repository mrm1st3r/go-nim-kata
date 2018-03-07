package main

import (
	"github.com/go-redis/redis"
	"fmt"
	"github.com/satori/go.uuid"
	"errors"
	"github.com/mrm1st3r/go-nim-kata/game"
)

var conn *redis.Client

func init() {
	conn = redis.NewClient(&redis.Options {
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

}

// persist either left matches or winner
func PersistGame(game game.State) {
	var value interface{}
	if game.Winner == "" {
		value = game.MatchesLeft
	} else {
		value = game.Winner
	}
	conn.Set(fmt.Sprintf("nim:%s", game.ID), value, 0)
}

func LoadGame(gameId interface{}) (game.State, error) {
	gameUuid, _ := uuid.FromString(gameId.(string))

	value := conn.Get(fmt.Sprintf("nim:%s", gameUuid))

	if value.Err() != nil {
		return game.State{}, errors.New("game not found")
	}

	matchesLeft, err := value.Int64()

	if err != nil {
		return gameWithWinner(gameUuid, value.Val()), nil
	}

	return activeGame(gameUuid, matchesLeft), nil
}

func gameWithWinner(id uuid.UUID, winner string) game.State {
	return game.State{ID: id, MatchesLeft: 0, Winner: winner}
}

func activeGame(id uuid.UUID, matchesLeft int64) game.State {
	return game.State{ID: id, MatchesLeft: int(matchesLeft),Winner: ""}
}