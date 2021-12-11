package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

	var basinSizes []int
	for _, lowPoint := range heightMap.FindLowPoints() {
		basin := heightMap.GetDownwardFlowForPoint(lowPoint)
		fmt.Printf("basin: %v", basin)

		basinSizes = append(basinSizes, len(basin))
		fmt.Printf("basin size: %d\n\n", len(basin))
	}

	sort.Ints(basinSizes)

	result := 1
	for i := len(basinSizes) - 1; i >= len(basinSizes)-3; i-- {
		result *= basinSizes[i]
	}

	fmt.Println("result", result)
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
