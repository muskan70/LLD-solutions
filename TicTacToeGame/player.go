package main

var id = 0

type Player struct {
	userId int
	name   string
	Piece  PlayingPiece
}

func NewPlayer(nm string, p PlayingPiece) Player {
	id += 1
	return Player{userId: id, name: nm, Piece: p}
}
