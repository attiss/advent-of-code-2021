package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	inputFile = "input"
)

func main() {
	instructions, err := readInputFromFile(inputFile)
	if err != nil {
		log.Fatalf("failed to get input: %v", err)
	}

	totalErrorScore := 0
	for _, instruction := range instructions {
		totalErrorScore += instruction.CheckSyntax()
	}

	fmt.Println(totalErrorScore)
}

func readInputFromFile(filePath string) ([]Instruction, error) {
	var instructions []Instruction

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instructions = append(instructions, Instruction(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return instructions, nil
}
