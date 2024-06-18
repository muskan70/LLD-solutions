package main

import "fmt"

const (
	WON         = "won"
	TIE         = "tie"
	IN_PROGRESS = "inProgress"
	NOT_STARTED = "notStarted"
)

type Game struct {
	Players []Player
	Board   PlayBoard
	Status  string
}

func NewGame() Game {
	fmt.Println("Enter No.of Players")
	var noofPlayers int
	fmt.Scanln(&noofPlayers)
	var players []Player
	for i := 0; i < noofPlayers; i++ {
		fmt.Println("Enter name and play symbol of Player ", i+1, ":")
		var name, symbol string
		fmt.Scanln(&name, &symbol)
		players = append(players, NewPlayer(name, NewPlayingPiece(symbol)))
	}
	fmt.Println("Enter PlayBoard size")
	var boardSize int
	fmt.Scanln(&boardSize)
	board := NewPlayBoard(boardSize)
	return Game{Players: players, Board: board, Status: NOT_STARTED}
}

func (g *Game) StartGame() {
	g.Status = IN_PROGRESS
	idx := 0

	g.Board.ShowBoard()
	for g.Status == IN_PROGRESS {
		fmt.Println("Player-", g.Players[idx].name, " Enter row, column for your turn :")
		var i, j int
		fmt.Scanln(&i, &j)
		for !g.Board.FillCellNShow(i, j, g.Players[idx].Piece) {
			fmt.Scanln(&i, &j)
		}
		g.Status = g.Board.WinCheck(g.Players[idx].Piece)
		if g.Status == WON {
			fmt.Println("Game ended with winner:", g.Players[idx].name)
			break
		}
		if g.Status == TIE {
			fmt.Println("Game ended with TIE")
			break
		}
		idx++
		if idx == len(g.Players) {
			idx = 0
		}
	}
}
