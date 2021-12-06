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
	lanternFishes, err := readInputFromFile(inputFile)
	if err != nil {
		log.Fatalf("failed to get input: %v", err)
	}

	for day := 1; day <= 80; day++ {
		lanternFishes.Cycle()
	}

	fmt.Println("population after 80 days", len(lanternFishes))
}

func readInputFromFile(filePath string) (Swarm, error) {
	var lanternFishes Swarm

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		for _, fishTimer := range strings.Split(line, ",") {
			parsedFishTimer, err := strconv.Atoi(fishTimer)
			if err != nil {
				return nil, err
			}
			lanternFishes = append(lanternFishes, NewLanternFish(parsedFishTimer))
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lanternFishes, nil
}
