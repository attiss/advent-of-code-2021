package main

import (
	"sort"
)

type CrabSubmarine struct {
	Position int
}

type Swarm []CrabSubmarine

func (s Swarm) Len() int {
	return len(s)
}

func (s Swarm) Less(i, j int) bool {
	return s[i].Position > s[j].Position
}

func (s Swarm) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Swarm) MeanPosition() int {
	sort.Sort(s)

	meanPlace := len(s) / 2

	if s[meanPlace].Position%2 == 0 {
		return s[meanPlace].Position
	}

	return (s[meanPlace-1].Position + s[meanPlace].Position) / 2
}

func (s Swarm) MoveToPosition(position int) (cost int) {
	cost = 0

	for _, crabSubmbarine := range s {
		if crabSubmbarine.Position > position {
			cost += crabSubmbarine.Position - position
		} else if crabSubmbarine.Position < position {
			cost += position - crabSubmbarine.Position
		}
	}

	return cost
}
