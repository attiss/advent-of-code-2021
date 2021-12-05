package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

const (
	inputFile = "input"
)

func main() {
	ventureLines, err := readInputFromFile(inputFile)
	if err != nil {
		log.Fatalf("failed to get input: %v", err)
	}

	max := 0
	for _, ventureLine := range ventureLines {
		if ventureLine.x1 > max {
			max = ventureLine.x1
		}
		if ventureLine.x2 > max {
			max = ventureLine.x2
		}
		if ventureLine.y1 > max {
			max = ventureLine.y1
		}
		if ventureLine.y2 > max {
			max = ventureLine.y2
		}
	}

	ventureMap := NewVentureMap(max)

	for _, ventureLine := range ventureLines {
		ventureMap.DrawVentureLine(ventureLine)
	}

	overlaps := ventureMap.CountOverlaps()

	fmt.Println("overlaps", overlaps)
}

func readInputFromFile(filePath string) ([]VentureLine, error) {
	var ventureLines []VentureLine

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineParser := regexp.MustCompile(`^([0-9]+),([0-9]+) -> ([0-9]+),([0-9]+)$`)
	for scanner.Scan() {
		line := scanner.Text()
		values := lineParser.FindStringSubmatch(line)

		if len(values) != 5 {
			return nil, fmt.Errorf("invalid input: %s", line)
		}

		x1, err := strconv.Atoi(values[1])
		if err != nil {
			return nil, err
		}

		y1, err := strconv.Atoi(values[2])
		if err != nil {
			return nil, err
		}

		x2, err := strconv.Atoi(values[3])
		if err != nil {
			return nil, err
		}

		y2, err := strconv.Atoi(values[4])
		if err != nil {
			return nil, err
		}

		ventureLines = append(ventureLines, VentureLine{x1: x1, y1: y1, x2: x2, y2: y2})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return ventureLines, nil
}
