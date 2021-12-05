package main

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
