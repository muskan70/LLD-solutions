package jumperFactory

import (
	"errors"
	"fmt"
	"snakesNLadders/constants"
	"snakesNLadders/player"
	"strconv"
)

type Ladder struct {
	Jumper
}

func NewLadder(start int, end int) (IJumper, error) {
	ladder := &Ladder{Jumper{
		StartPoint: start,
		EndPoint:   end},
	}
	if err := ladder.IsValid(); err != nil {
		return nil, err
	}
	return ladder, nil
}

func (l Ladder) IsValid() error {
	if l.StartPoint < l.EndPoint {
		return nil
	}
	return errors.New("Invalid Ladders entry : " + strconv.Itoa(l.StartPoint) + "," + strconv.Itoa(l.EndPoint))
}

func (l Ladder) GetJumperType() int {
	return constants.JTYPE_LADDER
}

func (l Ladder) Move(p *player.Player) {
	fmt.Println(p.GetName(), "climed a ladder at", l.StartPoint, "and moved from", p.GetCurrentPosition(), "to", l.EndPoint)
	p.Move(l.EndPoint)
}
