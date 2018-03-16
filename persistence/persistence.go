package persistence

import (
	"github.com/go-redis/redis"
	"fmt"
	"github.com/satori/go.uuid"
	"errors"
	"github.com/mrm1st3r/go-nim-kata/game"
	"strconv"
)

var conn *redis.Client

func init() {
	conn = redis.NewClient(&redis.Options {
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
}

func PersistGame(game game.State) {
	gameMap := make(map[string]interface{})
	gameMap["matchesLeft"] = game.MatchesLeft
	gameMap["winner"] = game.Winner

	conn.HMSet(dbKey(game.ID), gameMap)
}

func LoadGame(gameId interface{}) (game.State, error) {
	gameUuid, _ := uuid.FromString(gameId.(string))

	value := conn.HGetAll(dbKey(gameUuid))

	if value.Err() != nil {
		return game.State{}, errors.New("game not found")
	}

	m := value.Val()["matchesLeft"]
	matchesLeft, _ := strconv.Atoi(m)
	winner := value.Val()["winner"]

	return game.State{gameUuid, matchesLeft, winner}, nil
}

func dbKey(id uuid.UUID) string {
	return fmt.Sprintf("nim:%s", id)
}
