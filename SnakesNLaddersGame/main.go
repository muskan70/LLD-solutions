package main

import "snakesNLadders/game"

func main() {
	g := game.InitiateGame()
	if g != nil {
		g.StartGame()
	}
}
