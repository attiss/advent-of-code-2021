package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	inputFile = "input"
)

func main() {
	displays, err := readInputFromFile(inputFile)
	if err != nil {
		log.Fatalf("failed to get input: %v", err)
	}

	outputSum := 0
	for _, display := range displays {
		display.DecodeSignals()

		output, err := display.DecodeOutput()
		if err != nil {
			log.Fatalf("failed to decode output: %v", err)
		}

		outputSum += output[0]*1000 + output[1]*100 + output[2]*10 + output[3]
	}

	fmt.Println("total sum of outputs", outputSum)
}

func readInputFromFile(filePath string) ([]Display, error) {
	var displays []Display

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		lineParts := strings.Split(line, " | ")
		if len(lineParts) != 2 {
			return nil, errors.New("invalid input")
		}

		uniqueSignals := strings.Split(lineParts[0], " ")
		if len(uniqueSignals) != 10 {
			return nil, errors.New("invalid input")
		}

		output := strings.Split(lineParts[1], " ")
		if len(output) != 4 {
			return nil, errors.New("invalid input")
		}

		displays = append(displays, NewDisplay(uniqueSignals, output))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return displays, nil
}
