package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	inputFile = "input"
)

type Board [][]int

func (board Board) RemoveChosenNumber(chosenNumber int) {
	for y, boardRow := range board {
		for x, number := range boardRow {
			if number == chosenNumber {
				board[y][x] = -1
			}
		}
	}
}

func (board Board) IsWinningBoard() bool {
	for y := 0; y < len(board); y++ {
		isWinnerRow := true
		for x := 0; x < len(board[0]); x++ {
			if board[y][x] != -1 {
				isWinnerRow = false
				break
			}
		}
		if isWinnerRow {
			return true
		}
	}

	for x := 0; x < len(board[0]); x++ {
		isWinnerColumn := true
		for y := 0; y < len(board); y++ {
			if board[y][x] != -1 {
				isWinnerColumn = false
				break
			}
		}
		if isWinnerColumn {
			return true
		}
	}

	return false
}

func (board Board) GetScore() int {
	sum := 0

	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[0]); x++ {
			if board[y][x] != -1 {
				sum += board[y][x]
			}
		}
	}

	return sum
}

func main() {
	boards, choseNumbers, err := readInputFromFile(inputFile)
	if err != nil {
		log.Fatalf("failed to get input: %v", err)
	}

	for _, chosenNumber := range choseNumbers {
		for _, board := range boards {
			board.RemoveChosenNumber(chosenNumber)

			if board.IsWinningBoard() {
				fmt.Printf("winner board: %#v\n", board)
				fmt.Printf("winner board score: %d\n", board.GetScore())
				fmt.Printf("called number: %d\n", chosenNumber)
				fmt.Printf("result: %d\n", board.GetScore()*chosenNumber)
				return
			}
		}
	}
}

func readInputFromFile(filePath string) ([]Board, []int, error) {
	var boards []Board
	var chosenNumbers []int

	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var tmpBoard Board

	lineCount := 0
	scanner := bufio.NewScanner(file)
	whiteSpaceNormalizer := regexp.MustCompile(`\s+`)
	for scanner.Scan() {
		line := strings.TrimSpace(whiteSpaceNormalizer.ReplaceAllString(scanner.Text(), " "))
		lineCount++

		if lineCount == 1 {
			chosenNumbersString := strings.Split(line, ",")
			for _, chosenNchosenNumberString := range chosenNumbersString {
				chosenNumber, err := strconv.Atoi(chosenNchosenNumberString)
				if err != nil {
					return nil, nil, err
				}
				chosenNumbers = append(chosenNumbers, chosenNumber)
			}
			continue
		} else if lineCount == 2 {
			continue
		}

		if line == "" {
			boards = append(boards, tmpBoard)
			tmpBoard = nil
			continue
		}

		var boardRow []int
		lineValues := strings.Split(line, " ")
		for _, lineValue := range lineValues {
			value, err := strconv.Atoi(lineValue)
			if err != nil {
				return nil, nil, err
			}
			boardRow = append(boardRow, value)
		}

		tmpBoard = append(tmpBoard, boardRow)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return boards, chosenNumbers, nil
}
