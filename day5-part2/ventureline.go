package main

import "math"

type VentureLine struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func (vl VentureLine) IsHorizontal() bool {
	return vl.y1 == vl.y2
}

func (vl VentureLine) IsVertical() bool {
	return vl.x1 == vl.x2
}

func (vl VentureLine) IsDiagonal() bool {
	deltaX := math.Abs(float64(vl.x2 - vl.x1))
	deltaY := math.Abs(float64(vl.y2 - vl.y1))

	return deltaX == deltaY
}
