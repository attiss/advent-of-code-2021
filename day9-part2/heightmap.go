package main

import "fmt"

type Coordinate struct {
	X int
	Y int
}

func (c Coordinate) String() string {
	return fmt.Sprintf("(x=%d; y=%d)", c.X, c.Y)
}

type HeightMap [][]int

func (hm HeightMap) FindLowPoints() []Coordinate {
	var lowPoints []Coordinate

	isLowPoint := func(y, x int) bool {
		if y > 0 && hm[y-1][x] <= hm[y][x] {
			return false
		}
		if y < len(hm)-1 && hm[y+1][x] <= hm[y][x] {
			return false
		}
		if x > 0 && hm[y][x-1] <= hm[y][x] {
			return false
		}
		if x < len(hm[0])-1 && hm[y][x+1] <= hm[y][x] {
			return false
		}
		return true
	}

	for y := 0; y < len(hm); y++ {
		for x := 0; x < len(hm[0]); x++ {
			if isLowPoint(y, x) {
				lowPoints = append(lowPoints, Coordinate{Y: y, X: x})
			}
		}
	}

	return lowPoints
}

func (hm HeightMap) GetHeightAtPoint(point Coordinate) int {
	return hm[point.Y][point.X]
}

func (hm HeightMap) GetDownwardFlowForPoint(point Coordinate) []Coordinate {
	fmt.Println("visiting new point", point)
	downwardFlowPoints := []Coordinate{point}

	addPoints := func(points ...Coordinate) {
		for _, p := range points {
			alreadyPresent := false
			for _, downwardFlowPoint := range downwardFlowPoints {
				if downwardFlowPoint.Y == p.Y && downwardFlowPoint.X == p.X {
					alreadyPresent = true
					break
				}
			}
			if !alreadyPresent {
				downwardFlowPoints = append(downwardFlowPoints, p)
			}
		}
	}

	examineNeighbor := func(neighbor Coordinate) {
		fmt.Println("checking neighbor", neighbor)
		neighborHeight := hm.GetHeightAtPoint(neighbor)
		fmt.Println("neighbor height", neighborHeight)
		if neighborHeight != 9 && neighborHeight > hm.GetHeightAtPoint(point) {
			fmt.Println("neighbor is part of the dawnward flow, checking its neigbors")
			addPoints(hm.GetDownwardFlowForPoint(neighbor)...)
		} else {
			fmt.Println("neighbor is not part of the dawnward flow")
		}
		fmt.Println("finished checking neighbor", neighbor, "of", point)
	}

	if point.Y > 0 {
		examineNeighbor(Coordinate{Y: point.Y - 1, X: point.X})
	}
	if point.Y < len(hm)-1 {
		examineNeighbor(Coordinate{Y: point.Y + 1, X: point.X})
	}
	if point.X > 0 {
		examineNeighbor(Coordinate{Y: point.Y, X: point.X - 1})
	}
	if point.X < len(hm[0])-1 {
		examineNeighbor(Coordinate{Y: point.Y, X: point.X + 1})
	}

	fmt.Println("finished checking all neighbors", point)

	return downwardFlowPoints
}
