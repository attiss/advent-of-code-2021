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

	mostCommonBits := make([]int, len(input[0]))
	leastCommonBits := make([]int, len(input[0]))
	for x := 0; x < len(input[0]); x++ {
		column := getColumnFromMatrix(input, x)
		mostCommonBit := getMostCommonBit(column)

		mostCommonBits[x] = mostCommonBit

		if mostCommonBit == 1 {
			leastCommonBits[x] = 0
		} else {
			leastCommonBits[x] = 1
		}
	}

	decimalGammaRate := binaryToDecimal(mostCommonBits)
	decimalEpsilonRate := binaryToDecimal(leastCommonBits)

	fmt.Printf("decimal gamma rate: %d; decimal epsilon rate: %d; product: %d\n", decimalGammaRate, decimalEpsilonRate, decimalGammaRate*decimalEpsilonRate)
}

func getMostCommonBit(bits []int) int {
	sum := 0
	for _, b := range bits {
		sum += b
	}

	if sum >= len(bits)/2 {
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
