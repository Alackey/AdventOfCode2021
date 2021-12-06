package day4

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

func t_parseInput(t *testing.T) ([]string, []*board) {
	file, err := os.Open("input.txt")
	if err != nil {
		t.Errorf("Error opening file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	nums := scanner.Text()
	drawNumbers := strings.Split(nums, ",")

	var boards []*board
	for scanner.Scan() {
		// fmt.Println("Text: ", scanner.Text())
		if scanner.Text() == "" {
			// fmt.Println("----------------Blank Line----------------")
			var lines [][]string
			boards = append(boards, &board{lines})
		} else {
			line := strings.Split(strings.TrimSpace(scanner.Text()), " ")
			// fmt.Println(len(boards))
			// fmt.Println(line)
			curBoard := boards[len(boards)-1]
			curBoard.lines = append(curBoard.lines, line)
		}
	}
	return drawNumbers, boards
}

func TestPart1(t *testing.T) {
	expected := 33462
	actual, err := Part1(t_parseInput(t))

	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
