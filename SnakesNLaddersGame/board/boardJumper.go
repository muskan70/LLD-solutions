package board

import (
	"errors"
	constants "snakesNLadders/constants"
	jumper "snakesNLadders/jumperFactory"
	"strconv"
)

func (b *Board) IsValidBoardJumperCell(pos int) error {
	if pos == 1 || pos == b.Size*b.Size {
		errors.New("Jumper cannot be placed at starting and ending positions of board : " + strconv.Itoa(pos))
	}
	_, ok := b.Jumpers[pos]
	if ok {
		errors.New("Jumper Already exists at this cell : " + strconv.Itoa(pos))
	}
	return nil
}

func (b *Board) AddSnakes(noofSnakes int, snakes []jumper.Jumper) error {
	if noofSnakes != len(snakes) {
		return errors.New("Invalid NumberOfSnakes entry")
	}
	for i := 0; i < len(snakes); i++ {
		if err := b.IsValidCell(snakes[i].StartPoint); err != nil {
			return errors.Join(errors.New("Invalid snakes entry:"), err)
		}
		if err := b.IsValidCell(snakes[i].EndPoint); err != nil {
			return errors.Join(errors.New("Invalid snakes entry:"), err)
		}
		if err := b.IsValidBoardJumperCell(snakes[i].StartPoint); err != nil {
			return errors.Join(errors.New("Invalid snakes entry:"), err)
		}
		snake, err := jumper.JumperFactory(snakes[i].StartPoint, snakes[i].EndPoint, constants.JTYPE_SNAKE)
		if err != nil {
			return err
		} else {
			b.Jumpers[snakes[i].StartPoint] = snake
		}
	}
	return nil
}

func (b *Board) AddLadders(noofLadders int, ladders []jumper.Jumper) error {
	if noofLadders != len(ladders) {
		return errors.New("Invalid NumberOfLadders entry")
	}
	for i := 0; i < len(ladders); i++ {
		if err := b.IsValidCell(ladders[i].StartPoint); err != nil {
			return errors.Join(errors.New("Invalid ladders entry:"), err)
		}
		if err := b.IsValidCell(ladders[i].EndPoint); err != nil {
			return errors.Join(errors.New("Invalid ladders entry:"), err)
		}
		if err := b.IsValidBoardJumperCell(ladders[i].StartPoint); err != nil {
			return errors.Join(errors.New("Invalid ladders entry:"), err)
		}
		ladder, err := jumper.JumperFactory(ladders[i].StartPoint, ladders[i].EndPoint, constants.JTYPE_LADDER)
		if err != nil {
			return err
		} else {
			b.Jumpers[ladders[i].StartPoint] = ladder
		}
	}
	return nil
}

func (b *Board) AddCrocodiles(noofCrocs int, crocodiles []int) error {
	if noofCrocs != len(crocodiles) {
		return errors.New("Invalid NumberOfCrocodiles entry")
	}
	for i := 0; i < len(crocodiles); i++ {
		if err := b.IsValidCell(crocodiles[i]); err != nil {
			return errors.Join(errors.New("Invalid crocodiles entry:"), err)
		}
		if err := b.IsValidBoardJumperCell(crocodiles[i]); err != nil {
			return errors.Join(errors.New("Invalid crocodiles entry:"), err)
		}
		croc, err := jumper.JumperFactory(crocodiles[i], -1, constants.JTYPE_CROCODILE)
		if err != nil {
			return err
		} else {
			b.Jumpers[crocodiles[i]] = croc
		}
	}
	return nil
}

func (b *Board) AddMines(noofMines int, mines []int) error {
	if noofMines != len(mines) {
		return errors.New("Invalid NumberOfMines entry")
	}
	for i := 0; i < len(mines); i++ {
		if err := b.IsValidCell(mines[i]); err != nil {
			return errors.Join(errors.New("Invalid mines entry:"), err)
		}
		if err := b.IsValidBoardJumperCell(mines[i]); err != nil {
			return errors.Join(errors.New("Invalid mines entry:"), err)
		}
		mine, err := jumper.JumperFactory(mines[i], -1, constants.JTYPE_MINE)
		if err != nil {
			return err
		} else {
			b.Jumpers[mines[i]] = mine
		}
	}
	return nil
}
