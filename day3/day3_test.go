package day3

import (
	"bufio"
	"os"
	"testing"
)

func t_parseInput(t *testing.T) []string {
	file, err := os.Open("input.txt")
	if err != nil {
		t.Errorf("Error opening file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func TestPart1(t *testing.T) {
	expected := int64(3959450)
	actual, err := Part1(t_parseInput(t))

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	expected := int64(7440311)
	actual, err := Part2(t_parseInput(t))

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
