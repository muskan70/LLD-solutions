package jumperFactory

import (
	"errors"
	"fmt"
	"snakesNLadders/constants"
	"snakesNLadders/player"
	"strconv"
)

type Crocodile struct {
	Jumper
}

func NewCrocodile(start int) (IJumper, error) {
	crocodile := &Crocodile{Jumper{
		StartPoint: start,
		EndPoint:   start - 5},
	}
	if err := crocodile.IsValid(); err != nil {
		return nil, err
	}
	return crocodile, nil
}

func (c *Crocodile) IsValid() error {
	if c.EndPoint >= 1 {
		return nil
	}
	return errors.New("Invalid Crocodiles entry : " + strconv.Itoa(c.StartPoint))
}

func (c *Crocodile) GetJumperType() int {
	return constants.JTYPE_CROCODILE
}

func (c *Crocodile) Move(p *player.Player) {
	fmt.Println(p.GetName(), "bitten by crocodile at", c.StartPoint, "and moved from", p.GetCurrentPosition(), "to", c.EndPoint)
	p.Move(c.EndPoint)
}
