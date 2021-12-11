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
	heightMap, err := readInputFromFile(inputFile)
	if err != nil {
		log.Fatalf("failed to get input: %v", err)
	}

	lowPoints := heightMap.FindLowPoints()

	risk := 0
	for _, lowPoint := range lowPoints {
		risk += lowPoint + 1
	}

	fmt.Println("risk:", risk)
}

func readInputFromFile(filePath string) (HeightMap, error) {
	var input HeightMap

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var heights []int
		stringHeights := strings.Split(scanner.Text(), "")

		for _, stringHeight := range stringHeights {
			height, err := strconv.Atoi(stringHeight)
			if err != nil {
				return nil, err
			}

			heights = append(heights, height)
		}
		input = append(input, heights)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return input, nil
}
