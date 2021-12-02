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

type Vector struct {
	Direction string
	Magnitude int
}

func main() {
	input, err := readInputFromFile(inputFile)
	if err != nil {
		log.Fatalf("failed to get input: %v", err)
	}

	horisontalPosition := 0
	depth := 0

	for _, vector := range input {
		switch vector.Direction {
		case "forward":
			horisontalPosition += vector.Magnitude
		case "down":
			depth += vector.Magnitude
		case "up":
			depth -= vector.Magnitude
		default:
			log.Fatalf("invalid direction: %s", vector.Direction)
		}
	}

	fmt.Printf("horisontal position: %d; depth: %d\n", horisontalPosition, depth)
	fmt.Printf("product: %d\n", horisontalPosition*depth)
}

func readInputFromFile(filePath string) ([]Vector, error) {
	var input []Vector

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")

		if len(line) != 2 {
			return nil, fmt.Errorf("invalid input: %v", line)
		}

		direction := line[0]
		magnitude, err := strconv.Atoi(line[1])
		if err != nil {
			return nil, err
		}

		input = append(input, Vector{Direction: direction, Magnitude: magnitude})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return input, nil
}
