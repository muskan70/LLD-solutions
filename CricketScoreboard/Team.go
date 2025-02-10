package main

type Team struct {
	TeamId                int
	TeamName              string
	Players               []Player
	CurrentBattingScore   BattingScore
	CurrentBowlingScore   BowlingScore
	OversCompleted        int
	CurrentStrike         int // Represent PlayerId
	CurrentBattingPlayers []int
	CurrentBowler         int // Represent PlayerId
}

func (t *Team) AddCurrentOverScore() {
	//Step1: loop over scores and add score to currentPlayer
	// if score is 1 or 3 change current player -> swap current strike
	// if ball is wide, no ball ,add to extra score of team
	// if wicket  ball then move to nextplayer, change status of current player and nextPlayer
	//Increment NoofOvers

	// In the end of over score update, change strike
}
