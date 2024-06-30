package game

import (
	"fmt"
	"snakesNLadders/board"
	"snakesNLadders/constants"
	"snakesNLadders/dice"
	"snakesNLadders/player"
)

type Game struct {
	Players player.Players
	Board   board.Board
	Dice    dice.Dice
	Status  int
}

func (g *Game) StartGame() {
	g.Status = constants.GAME_STATUS_IN_PROGRESS
	idx := 0

	for g.Status == constants.GAME_STATUS_IN_PROGRESS {
		blockedTurns := g.Players[idx].GetBlockedTurns()
		if blockedTurns != 0 {
			fmt.Println(g.Players[idx].GetName(), "Your turn is blocked because you are standing on mine")
			g.Players[idx].SetBlockTurns(blockedTurns - 1)
			idx++
			if idx == len(g.Players) {
				idx = 0
			}
			continue
		}

		diceVal := g.Dice.RollDice(g.Players[idx])
		g.Board.MovePlayer(g.Players[idx], diceVal)

		g.Players.CheckPositionOverlap(idx)
		g.Status = g.Board.CheckWinCondition(g.Players[idx])
		if g.Status == constants.GAME_STATUS_WON {
			fmt.Println("Game ended with winner:", g.Players[idx].GetName())
			break
		}
		idx++
		if idx == len(g.Players) {
			idx = 0
		}
	}
}
