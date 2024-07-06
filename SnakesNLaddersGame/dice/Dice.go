package dice

import (
	"fmt"
	"math/rand"
	"snakesNLadders/constants"
	"snakesNLadders/player"
)

type Dice struct {
	noofDice         int
	movementStrategy IMovementStrategy
	entry            string
}

func NewDice(noofDice int, strategy string, entry string) Dice {
	d := Dice{noofDice: noofDice, entry: entry}
	switch strategy {
	case constants.MOVEMENT_STRATEGY_SUM:
		d.movementStrategy = &SumStrategy{}
	case constants.MOVEMENT_STRATEGY_MIN:
		d.movementStrategy = &MinStrategy{}
	case constants.MOVEMENT_STRATEGY_MAX:
		d.movementStrategy = &MaxStrategy{}
	}
	return d
}

func (d *Dice) RollDice(p *player.Player) int {
	var vals []int

	if d.entry == constants.DICE_ENTRY_AUTOMATED {
		for i := 0; i < d.noofDice; i++ {
			vals = append(vals, rand.Intn(6)+1)
		}
	} else {
		fmt.Println(p.GetName(), "Enter", d.noofDice, "dice values separated by next line:")
		var val int
		for i := 0; i < d.noofDice; i++ {
			fmt.Scanln(&val)
			for val < 1 || val > 6 {
				fmt.Println("Invalid dice entry")
				fmt.Scanln(&val)
			}
			vals = append(vals, val)
		}
	}
	diceVal := d.getDiceValueWRTMovementStrategy(vals)
	fmt.Println(p.GetName(), "rolled his dice and got", vals, "so moves =", diceVal)
	return diceVal
}

func (d *Dice) getDiceValueWRTMovementStrategy(vals []int) int {
	if d.noofDice == 1 {
		return vals[0]
	}
	return d.movementStrategy.GetMoves(vals)
}
