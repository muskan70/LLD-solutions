package models

import "strconv"

type Seat struct {
	Id           string
	RowNo        int
	ColNo        int
	SeatCategory int
}

func NewSeat(rowNo, colNo, category int) *Seat {
	return &Seat{
		Id:           GetRowString(rowNo) + strconv.Itoa(colNo),
		RowNo:        rowNo,
		ColNo:        colNo,
		SeatCategory: category,
	}
}

func GetRowString(rowNo int) string {
	rowId := ""
	if rowNo > 26 {
		d := rowNo / 26
		rowId += string(d + 64)
		rowNo = rowNo % 26
	}
	rowId += string(rowNo + 64)
	return rowId
}

func GetRowNo(rowId string) int {
	rowNo := 0
	for i := range rowId {
		rowNo += int(rowId[i] - 64)
	}
	return rowNo
}

func GetSeatRowNCol(id string) (int, int) {
	rowNo := 0
	colNo := 0
	for i := range id {
		if id[i] >= 'A' && id[i] <= 'Z' {
			continue
		} else {
			rowNo = GetRowNo(id[:i])
			colNo, _ = strconv.Atoi(id[i:])
			break
		}
	}
	return rowNo, colNo
}
