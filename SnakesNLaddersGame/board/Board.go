package board

import (
	"errors"
	"fmt"
	"snakesNLadders/constants"
	jumper "snakesNLadders/jumperFactory"
	"snakesNLadders/player"
	"strconv"
)

type Board struct {
	Size            int // Board ->(size *size)
	Jumpers         map[int]jumper.IJumper
	PlayersPosition map[int]int
}

func NewBoard(sz int) Board {
	plBrd := Board{
		Size:            sz,
		Jumpers:         make(map[int]jumper.IJumper),
		PlayersPosition: make(map[int]int),
	}
	return plBrd
}

func (b *Board) IsValidCell(pos int) error {
	if pos >= 1 && pos <= b.Size*b.Size {
		return nil
	}
	return errors.New("Invalid board cell : " + strconv.Itoa(pos))
}

func (b *Board) MovePlayer(idx int, players player.Players, diceVal int) {
	currPos := players[idx].GetCurrentPosition()
	nextPos := currPos + diceVal
	if nextPos > b.Size*b.Size {
		fmt.Println(players[idx].GetName(), "cannot move as dice value exceeds board size")
		return
	}
	fmt.Println(players[idx].GetName(), "moved from", currPos, "to", nextPos)
	players[idx].Move(nextPos)
	delete(b.PlayersPosition, currPos)

	jObj, ok := b.Jumpers[nextPos]
	for ok {
		jObj.Move(players[idx])
		nextPos = players[idx].GetCurrentPosition()
		if jObj.GetJumperType() == constants.JTYPE_MINE {
			break
		}
		jObj, ok = b.Jumpers[nextPos]
	}

	//Check position Overlap
	currPlayer, ok := b.PlayersPosition[nextPos]
	if ok {
		fmt.Println(players[currPlayer].GetName(), "moved from", nextPos, "to 0 as", players[idx].GetName(), "occupied this position")
		players[currPlayer].Move(0)
	}
	b.PlayersPosition[nextPos] = idx
}

func (b *Board) CheckWinCondition(p *player.Player) int {
	if p.GetCurrentPosition() == b.Size*b.Size {
		return constants.GAME_STATUS_WON
	}
	return constants.GAME_STATUS_IN_PROGRESS
}
