package jumperFactory

import "snakesNLadders/constants"

func JumperFactory(start int, end int, jumperType int) (IJumper, error) {
	switch jumperType {
	case constants.JTYPE_SNAKE:
		return NewSnake(start, end)
	case constants.JTYPE_LADDER:
		return NewLadder(start, end)
	case constants.JTYPE_CROCODILE:
		return NewCrocodile(start)
	case constants.JTYPE_MINE:
		return NewMine(start)
	}
	return nil, nil
}
