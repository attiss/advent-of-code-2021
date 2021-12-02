package main

import (
	"bufio"
	"fmt"
	"log"
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

	increases := 0

	for i := 1; i < len(input); i = i + 3 {
		if input[i] > input[i-1] {
			increases++
		}
	}

	fmt.Printf("increases in input: %d\n", increases)
}

func readInputFromFile(filePath string) (input []int, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var parsedInt int
		parsedInt, err = strconv.Atoi(scanner.Text())
		input = append(input, parsedInt)
	}

	if err = scanner.Err(); err != nil {
		return
	}

	return
}
