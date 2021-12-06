package main

import "fmt"

type Swarm map[int]int

func (s Swarm) Cycle() {
	var fishSpawns int

	for daysToSpawn := 0; daysToSpawn <= 8; daysToSpawn++ {
		if population, ok := s[daysToSpawn]; ok {
			if daysToSpawn == 0 {
				fishSpawns = population
			} else {
				s[daysToSpawn-1] = population
			}
			s[daysToSpawn] = 0
		}
	}

	s[6] += fishSpawns
	s[8] = fishSpawns
}

func (s Swarm) Population() int {
	totalPopulation := 0
	for _, populationCount := range s {
		totalPopulation += populationCount
	}
	return totalPopulation
}

func (s Swarm) Print() {
	for daysToSpawn := 0; daysToSpawn <= 8; daysToSpawn++ {
		if populationCount, ok := s[daysToSpawn]; ok {
			fmt.Println("days to spawn:", daysToSpawn, "population:", populationCount)
		}
	}
}
