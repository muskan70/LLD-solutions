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
	Size    int // Board ->(size *size)
	Jumpers map[int]jumper.IJumper
}

func NewBoard(sz int) Board {
	plBrd := Board{
		Size:    sz,
		Jumpers: make(map[int]jumper.IJumper),
	}
	return plBrd
}

func (b *Board) IsValidCell(pos int) error {
	if pos >= 1 && pos <= b.Size*b.Size {
		return nil
	}
	return errors.New("Invalid board cell : " + strconv.Itoa(pos))
}

func (b *Board) MovePlayer(p *player.Player, diceVal int) {
	nextPos := p.GetCurrentPosition() + diceVal
	if nextPos > b.Size*b.Size {
		fmt.Println(p.GetName(), "cannot move as dice value exceeds board size")
		return
	}
	fmt.Println(p.GetName(), "moved from", p.GetCurrentPosition(), "to", nextPos)
	p.Move(nextPos)

	jObj, ok := b.Jumpers[nextPos]
	for ok {
		jObj.Move(p)
		nextPos = p.GetCurrentPosition()
		if jObj.GetJumperType() == constants.JTYPE_MINE {
			break
		}
		jObj, ok = b.Jumpers[nextPos]
	}
}

func (b *Board) CheckWinCondition(p *player.Player) int {
	if p.GetCurrentPosition() == b.Size*b.Size {
		return constants.GAME_STATUS_WON
	}
	return constants.GAME_STATUS_IN_PROGRESS
}
