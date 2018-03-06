package main

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"strconv"
	"fmt"
	"encoding/json"
	"nim/game"
)

func main() {
	router := fasthttprouter.New()
	router.GET("/game/:gameId", GetGameState)
	router.POST("/game", StartGame)
	router.POST("/game/:gameId", PlayGame)
	fasthttp.ListenAndServe(":8080", router.Handler)
}

func StartGame(ctx *fasthttp.RequestCtx) {
	g := game.New()
	PersistGame(g)
	play(ctx, g)
}


func PlayGame(ctx *fasthttp.RequestCtx) {
	g, err := LoadGame(ctx.UserValue("gameId"))
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
	PersistGame(g)

	//ctx.Redirect(fmt.Sprintf("/game/%s",g.ID), 302)
	marshaledGame, _ := json.Marshal(g)
	fmt.Fprintf(ctx, string(marshaledGame))
}

func GetGameState(ctx *fasthttp.RequestCtx) {
	g, err := LoadGame(ctx.UserValue("gameId"))

	if err != nil {
		ctx.Error(err.Error(), 404)
		return
	}

	marshaledGame, _ := json.Marshal(g)
	fmt.Fprintf(ctx, string(marshaledGame))
}
