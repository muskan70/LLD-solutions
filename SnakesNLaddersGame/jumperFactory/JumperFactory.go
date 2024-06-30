package jumperFactory

import "snakesNLadders/constants"

func JumperFactory(start int, end int, jumperType int) (IJumper, error) {
	if jumperType == constants.JTYPE_SNAKE {
		return NewSnake(start, end)
	} else if jumperType == constants.JTYPE_LADDER {
		return NewLadder(start, end)
	} else if jumperType == constants.JTYPE_CROCODILE {
		return NewCrocodile(start)
	} else if jumperType == constants.JTYPE_MINE {
		return NewMine(start)
	}
	return nil, nil
}
