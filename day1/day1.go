package day1

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func getInput() (*bufio.Scanner, error) {
	file, err := os.Open("day1/input.txt")
	if err != nil {
		return nil, err
	}
	return bufio.NewScanner(file), nil
}

func Part1() (int, error) {
	scanner, err := getInput()
	if err != nil {
		return 0, err
	}

	var count int
	prev := math.MaxInt

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return 0, err
		}

		if num > prev {
			count++
		}
		prev = num
	}

	return count, nil
}

func Part2() (int, error) {
	scanner, err := getInput()
	if err != nil {
		return 0, err
	}

	var nums []int
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return 0, err
		}
		nums = append(nums, num)
	}

	var sums []int
	for i := 0; i < len(nums)-2; i++ {
		sum := nums[i] + nums[i+1] + nums[i+2]
		sums = append(sums, sum)
	}

	var count int
	prev := math.MaxInt
	for _, v := range sums {
		if v > prev {
			count++
		}
		prev = v
	}
	return count, nil
}
