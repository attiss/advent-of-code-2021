package main

type CrabSubmarine struct {
	Position int
}

func (cs CrabSubmarine) CostOfMoving(positions int) int {
	if positions == 0 {
		return 0
	}
	return cs.CostOfMoving(positions-1) + positions
}

type Swarm []CrabSubmarine

func (s Swarm) MoveToPosition(position int) (cost int) {
	cost = 0

	for _, crabSubmbarine := range s {
		if crabSubmbarine.Position > position {
			cost += crabSubmbarine.CostOfMoving(crabSubmbarine.Position - position)
		} else if crabSubmbarine.Position < position {
			cost += crabSubmbarine.CostOfMoving(position - crabSubmbarine.Position)
		}
	}

	return cost
}

func (s Swarm) AveragePosition() int {
	sum := 0
	count := 0

	for _, crabSubmarine := range s {
		sum += crabSubmarine.Position
		count++
	}

	return sum / count
}
