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

	for day := 1; day <= 256; day++ {
		fmt.Println("day", day)
		lanternFishes.Print()
		lanternFishes.Cycle()
	}

	fmt.Println("population after 256 days", lanternFishes.Population())
}

func readInputFromFile(filePath string) (Swarm, error) {
	lanternFishes := make(Swarm)

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		for _, spawnTimer := range strings.Split(line, ",") {
			parsedSpawnTimer, err := strconv.Atoi(spawnTimer)
			if err != nil {
				return nil, err
			}
			lanternFishes[parsedSpawnTimer]++
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lanternFishes, nil
}
