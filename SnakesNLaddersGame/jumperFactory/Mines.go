package jumperFactory

import (
	"fmt"
	"snakesNLadders/constants"
	"snakesNLadders/player"
)

type Mine struct {
	Jumper
}

func NewMine(start int) (IJumper, error) {
	mine := &Mine{Jumper{
		StartPoint: start,
		EndPoint:   start},
	}
	return mine, nil
}

func (m Mine) IsValid() error {
	return nil
}

func (m Mine) Move(p *player.Player) {
	fmt.Println(p.GetName(), "fell into mine at", m.StartPoint, "and his next two turns are blocked")
	p.SetBlockTurns(2)
}

func (m Mine) GetJumperType() int {
	return constants.JTYPE_MINE
}
