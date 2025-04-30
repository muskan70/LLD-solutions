package domain

import "math"

type Location struct {
	XCoord float64
	YCoord float64
}

func NewLocation(x, y float64) Location {
	return Location{
		XCoord: x,
		YCoord: y,
	}
}

func GetDistance(l1, l2 Location) float64 {
	x := l1.XCoord - l2.XCoord
	y := l1.YCoord - l2.YCoord
	return math.Sqrt((x * x) + (y * y))
}
