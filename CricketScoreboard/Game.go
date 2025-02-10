package main

type Game struct {
	GameId                int
	NoofPlayersInEachTeam int
	Teams                 []Team
	TotalOvers            int
	BattingOrder          int // teamId who bats first
	TeamSize              int
	PlayersBattingOrder   []int
}
