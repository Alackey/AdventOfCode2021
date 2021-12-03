package day2

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	file, err := os.Open("input.txt")
	if err != nil {
		t.Errorf("Error opening file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var input []*line
	for scanner.Scan() {
		lineSplit := strings.Split(scanner.Text(), " ")

		num, err := strconv.Atoi(lineSplit[1])
		if err != nil {
			t.Errorf("Error converting %s to int", lineSplit[1])
		}
		input = append(input, &line{lineSplit[0], num})
	}

	expected := 1746616
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

	var input []*line
	for scanner.Scan() {
		lineSplit := strings.Split(scanner.Text(), " ")

		num, err := strconv.Atoi(lineSplit[1])
		if err != nil {
			t.Errorf("Error converting %s to int", lineSplit[1])
		}
		input = append(input, &line{lineSplit[0], num})
	}

	expected := 1741971043
	actual, err := Part2(input)

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
