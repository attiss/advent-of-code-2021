package main

import (
	"fmt"
	"math"
)

type VentureMap [][]int

func NewVentureMap(size int) VentureMap {
	vm := make(VentureMap, size+1)

	for row := 0; row < size+1; row++ {
		vm[row] = make([]int, size+1)
	}

	return vm
}

func (vm VentureMap) DrawVentureLine(vl VentureLine) {
	if vl.IsHorizontal() {
		start := vl.x1
		end := vl.x2
		if vl.x1 > vl.x2 {
			start = vl.x2
			end = vl.x1
		}

		for x := start; x <= end; x++ {
			vm[vl.y1][x]++
		}
		return
	}

	if vl.IsVertical() {
		start := vl.y1
		end := vl.y2
		if vl.y1 > vl.y2 {
			start = vl.y2
			end = vl.y1
		}

		for y := start; y <= end; y++ {
			vm[y][vl.x1]++
		}
		return
	}

	if vl.IsDiagonal() {
		deltaX := vl.x2 - vl.x1
		deltaY := vl.y2 - vl.y1

		for i := 0; i <= int(math.Abs(float64(deltaX))); i++ {
			x := vl.x1 + i
			if deltaX < 0 {
				x = vl.x1 - i
			}
			y := vl.y1 + i
			if deltaY < 0 {
				y = vl.y1 - i
			}
			vm[y][x]++
		}

		return
	}

	fmt.Println("venture line is not horizontal nor vertical nor diagonal, skipped", vl)
}

func (vm VentureMap) CountOverlaps() int {
	overlaps := 0
	for y := 0; y < len(vm); y++ {
		for x := 0; x < len(vm[0]); x++ {
			if vm[y][x] > 1 {
				overlaps++
			}
		}
	}

	return overlaps
}
