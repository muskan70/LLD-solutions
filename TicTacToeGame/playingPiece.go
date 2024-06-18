package main

type PlayingPiece struct {
	Symbol string
}

func NewPlayingPiece(s string) PlayingPiece {
	return PlayingPiece{Symbol: s}
}

func NewPlayingPieceUsingPlayerPiece(p PlayingPiece) PlayingPiece {
	return PlayingPiece{Symbol: p.Symbol}
}
