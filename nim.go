package main

import (
	"encoding/json"
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/mrm1st3r/go-nim-kata/game"
	"github.com/mrm1st3r/go-nim-kata/persistence"
	"github.com/valyala/fasthttp"
	"strconv"
)

func main() {
	router := fasthttprouter.New()
	router.GET("/game/:gameId", getGameState)
	router.POST("/game", startGame)
	router.POST("/game/:gameId", playGame)
	fasthttp.ListenAndServe(":8080", router.Handler)
}

func startGame(ctx *fasthttp.RequestCtx) {
	g := game.New()
	play(ctx, g)
}

func playGame(ctx *fasthttp.RequestCtx) {
	g, err := persistence.LoadGame(ctx.UserValue("gameId"))
	if err != nil {
		ctx.Error(err.Error(), 404)
		return
	}
	play(ctx, g)
}

func play(ctx *fasthttp.RequestCtx, g game.State) {
	number, e := strconv.Atoi(string(ctx.PostBody()))
	if e != nil {
		fmt.Fprintf(ctx, "Not a number: %s (%s)", ctx.PostBody(), e)
		ctx.SetStatusCode(400)
		return
	}

	g, err := game.Play(g, number)
	if err != nil {
		ctx.Error(err.Error(), 400)
		return
	}
	persistence.PersistGame(g)

	marshaledGame, _ := json.Marshal(g)
	fmt.Fprintf(ctx, string(marshaledGame))
}

func getGameState(ctx *fasthttp.RequestCtx) {
	g, err := persistence.LoadGame(ctx.UserValue("gameId"))

	if err != nil {
		ctx.Error(err.Error(), 404)
		return
	}

	marshaledGame, _ := json.Marshal(g)
	fmt.Fprintf(ctx, string(marshaledGame))
}
