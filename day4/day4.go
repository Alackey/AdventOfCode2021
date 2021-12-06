package day4

import (
	"errors"
	"strconv"
	"strings"
)

type board struct {
	lines [][]string
}

func sumOfUnmarked(board *board) (int, error) {
	var sum int
	for _, row := range board.lines {
		for _, cell := range row {
			if !strings.Contains(cell, "*") {
				num, err := strconv.Atoi(cell)
				if err != nil {
					return -1, err
				}
				sum += num
			}
		}
	}
	return sum, nil
}

func won(board *board) bool {
	// Check rows
	for i := 0; i < 5; i++ {
		if strings.Contains(board.lines[i][0], "*") &&
			strings.Contains(board.lines[i][1], "*") &&
			strings.Contains(board.lines[i][2], "*") &&
			strings.Contains(board.lines[i][3], "*") &&
			strings.Contains(board.lines[i][4], "*") {
			return true
		}
	}

	// Check columns
	for i := 0; i < 5; i++ {
		if strings.Contains(board.lines[0][i], "*") &&
			strings.Contains(board.lines[1][i], "*") &&
			strings.Contains(board.lines[2][i], "*") &&
			strings.Contains(board.lines[3][i], "*") &&
			strings.Contains(board.lines[4][i], "*") {
			return true
		}
	}
	return false
}

func markBoard(drawNumber string, board *board) {
	for i := 0; i < len(board.lines); i++ {
		for j := 0; j < len(board.lines[i]); j++ {
			if board.lines[i][j] == drawNumber {
				board.lines[i][j] = "*" + board.lines[i][j]
			}
		}
	}
}

// Part1 returns the bingo winner sum * drawn number
func Part1(drawNumbers []string, boards []*board) (int, error) {
	for i := 0; i < len(drawNumbers); i++ {
		for j := 0; j < len(boards); j++ {
			markBoard(drawNumbers[i], boards[j])

			won := won(boards[j])
			if won {
				sum, err := sumOfUnmarked(boards[j])
				if err != nil {
					return -1, err
				}

				drawNumber, err := strconv.Atoi(drawNumbers[i])
				if err != nil {
					return -1, err
				}
				return sum * drawNumber, nil
			}
		}
	}
	return -1, errors.New("no winning board found")
}
