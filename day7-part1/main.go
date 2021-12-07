package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	inputFile = "input"
)

func main() {
	crabSubmarines, err := readInputFromFile(inputFile)
	if err != nil {
		log.Fatalf("failed to get input: %v", err)
	}

	meanPosition := crabSubmarines.MeanPosition()
	totalCost := crabSubmarines.MoveToPosition(meanPosition)

	fmt.Println("mean position", meanPosition)
	fmt.Println("total cost to get to mean position", totalCost)
}

func readInputFromFile(filePath string) (Swarm, error) {
	var crabSubmarines Swarm

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		for _, positionString := range strings.Split(line, ",") {
			position, err := strconv.Atoi(positionString)
			if err != nil {
				return nil, err
			}
			crabSubmarines = append(crabSubmarines, CrabSubmarine{Position: position})
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return crabSubmarines, nil
}
