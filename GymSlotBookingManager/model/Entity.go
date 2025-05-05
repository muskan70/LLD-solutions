package model

import "math"

type Location struct {
	xcoord float64
	ycoord float64
}

func NewLocation(x, y float64) Location {
	return Location{
		xcoord: x,
		ycoord: y,
	}
}

func (l *Location) IsValid() bool {
	return l.xcoord != 0 && l.ycoord != 0
}

func GetDistance(l1, l2 Location) float64 {
	x := l1.xcoord - l2.xcoord
	y := l1.ycoord - l2.ycoord
	return math.Sqrt((x * x) + (y * y))
}

type Timing struct {
	StartTime int
	EndTime   int
}

func (t *Timing) IsValid() bool {
	if t.StartTime >= t.EndTime || t.StartTime < 0 || t.StartTime > 24 || t.EndTime < 0 || t.EndTime > 24 {
		return false
	}
	return true
}
