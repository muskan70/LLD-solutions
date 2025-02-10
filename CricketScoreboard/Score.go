package main

type BattingScore struct {
	TotalScore     int
	NoofSix        int
	NoofFour       int
	NoofBallsFaced int
}

type TeamLevelScore struct {
	ExtraScore int
}

type BowlingScore struct {
	WideBalls        int
	NoBalls          int
	NoofOvers        int
	TotalBalls       int
	NoofWicketsTaken int
	RunsConceded     int
}
