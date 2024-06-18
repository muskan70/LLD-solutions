package main

import "fmt"

type PlayBoard struct {
	Size        int
	Board       [][]PlayingPiece
	CellsFilled int
}

func NewPlayBoard(sz int) PlayBoard {
	plBrd := PlayBoard{Size: sz, CellsFilled: 0}
	plBrd.Board = make([][]PlayingPiece, sz)
	for i := range plBrd.Board {
		plBrd.Board[i] = make([]PlayingPiece, sz)
	}
	return plBrd
}

func (b *PlayBoard) ShowBoard() {
	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			fmt.Printf(b.Board[i][j].Symbol)
			if j != b.Size-1 {
				fmt.Printf(" | ")
			}
		}
		fmt.Println()
	}
}

func (b *PlayBoard) FillCellNShow(i int, j int, p PlayingPiece) bool {
	ep := PlayingPiece{}
	if b.Board[i][j] == ep {
		b.Board[i][j] = NewPlayingPieceUsingPlayerPiece(p)
		b.ShowBoard()
		return true
	}
	fmt.Println("Already Filled Cell, Please Choose empty cell")
	return false
}

func (b *PlayBoard) WinCheck(p PlayingPiece) string {
	b.CellsFilled++
	if b.CellsFilled == b.Size*b.Size {
		return TIE
	}
	rowFreq := make(map[PlayingPiece]int)
	colFreq := make(map[PlayingPiece]int)
	diaFreq := make(map[PlayingPiece]int)
	backDiaFreq := make(map[PlayingPiece]int)
	discardedRows := 0
	discardedCols := 0
	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			rowFreq[b.Board[i][j]]++
			colFreq[b.Board[j][i]]++
			if i == j {
				diaFreq[b.Board[i][j]]++
			}
			if i+j == b.Size-1 {
				backDiaFreq[b.Board[i][j]]++
			}
		}
		if rowFreq[p] == b.Size || colFreq[p] == b.Size {
			return WON
		}
		delete(rowFreq, PlayingPiece{})
		delete(colFreq, PlayingPiece{})
		if len(rowFreq) > 1 {
			discardedRows++
		}
		if len(colFreq) > 1 {
			discardedCols++
		}
		clear(rowFreq)
		clear(colFreq)
	}
	if diaFreq[p] == b.Size {
		return WON
	}
	if backDiaFreq[p] == b.Size {
		return WON
	}
	delete(diaFreq, PlayingPiece{})
	delete(backDiaFreq, PlayingPiece{})
	if len(diaFreq) > 1 && len(backDiaFreq) > 1 && discardedRows == b.Size && discardedCols == b.Size {
		return TIE
	}

	return IN_PROGRESS
}
