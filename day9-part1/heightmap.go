package main

type HeightMap [][]int

func (hm HeightMap) FindLowPoints() []int {
	var lowPoints []int

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
				lowPoints = append(lowPoints, hm[y][x])
			}
		}
	}

	return lowPoints
}
