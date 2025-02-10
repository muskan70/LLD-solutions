package main

type Player struct {
	PlayerId            int
	Name                string
	CurrentBattingScore BattingScore
	CurrentBowlingScore BowlingScore
	Status              string
}
