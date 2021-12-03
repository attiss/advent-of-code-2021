package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

const (
	inputFile = "input"
)

func main() {
	input, err := readInputFromFile(inputFile)
	if err != nil {
		log.Fatalf("failed to get input: %v", err)
	}

	oxygenGeneratorRate := findRate(input, true)
	co2ScrubberRate := findRate(input, false)

	fmt.Printf("oxygenGeneratorRate: %d; co2ScrubberRate: %d; product: %d\n", oxygenGeneratorRate, co2ScrubberRate, oxygenGeneratorRate*co2ScrubberRate)
}

func findRate(input [][]int, keepMostCommonBits bool) int {
	var rate int

	values := make([][]int, len(input))
	copy(values, input)
	for x := 0; x < len(input[0]); x++ {
		column := getColumnFromMatrix(values, x)
		fmt.Println("column", x, column)
		mostCommonBit := getMostCommonBit(column)
		fmt.Println("most common bit", mostCommonBit)

		var potentialValues [][]int
		for y := 0; y < len(values); y++ {
			if (mostCommonBit == values[y][x] && keepMostCommonBits) || (mostCommonBit != values[y][x] && !keepMostCommonBits) {
				potentialValues = append(potentialValues, values[y])
				fmt.Println("row", values[y], "kept")
			} else {
				fmt.Println("row", values[y], "dropped")
			}
		}

		if len(potentialValues) == 1 {
			rate = binaryToDecimal(potentialValues[0])
			break
		}

		values = potentialValues
	}

	return rate
}

func getMostCommonBit(bits []int) int {
	sum := 0
	for _, b := range bits {
		sum += b
	}

	if sum >= int(math.Ceil(float64(len(bits))/2)) {
		return 1
	}
	return 0
}

func getColumnFromMatrix(values [][]int, x int) []int {
	var column []int
	for y := 0; y < len(values); y++ {
		column = append(column, values[y][x])
	}
	return column
}

func binaryToDecimal(binary []int) int {
	value := 0
	for i, v := range binary {
		value += v * int(math.Pow(2, float64(len(binary)-i-1)))
	}
	return value
}

func readInputFromFile(filePath string) ([][]int, error) {
	var input [][]int

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		var lineValues []int

		for _, c := range line {
			value, err := strconv.Atoi(string(c))
			if err != nil {
				return nil, err
			}

			lineValues = append(lineValues, value)
		}

		input = append(input, lineValues)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return input, nil
}
