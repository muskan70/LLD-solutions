package jumperFactory

import (
	"errors"
	"fmt"
	"snakesNLadders/constants"
	"snakesNLadders/player"
	"strconv"
)

type Snake struct {
	Jumper
}

func NewSnake(start int, end int) (IJumper, error) {
	snake := &Snake{Jumper{
		StartPoint: start,
		EndPoint:   end},
	}
	if err := snake.IsValid(); err != nil {
		return nil, err
	}
	return snake, nil
}

func (s Snake) IsValid() error {
	if s.StartPoint > s.EndPoint {
		return nil
	}
	return errors.New("Invalid Snakes entry : " + strconv.Itoa(s.StartPoint) + "," + strconv.Itoa(s.EndPoint))
}

func (s Snake) Move(p *player.Player) {
	fmt.Println(p.GetName(), "is bitten by snake at", s.StartPoint, "and moved from", p.GetCurrentPosition(), "to", s.EndPoint)
	p.Move(s.EndPoint)
}

func (s Snake) GetJumperType() int {
	return constants.JTYPE_SNAKE
}
