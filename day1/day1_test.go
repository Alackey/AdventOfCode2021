package day1

import (
	"bufio"
	"os"
	"strconv"
	"testing"
)

func TestPart1(t *testing.T) {
	file, err := os.Open("input.txt")
	if err != nil {
		t.Errorf("Error opening file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var input []int
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			t.Errorf("Error converting %s to int", scanner.Text())
		}
		input = append(input, num)
	}

	expected := 1715
	actual, err := Part1(input)

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	file, err := os.Open("input.txt")
	if err != nil {
		t.Errorf("Error opening file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var input []int
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			t.Errorf("Error converting %s to int", scanner.Text())
		}
		input = append(input, num)
	}

	expected := 1739
	actual, err := Part2(input)

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
