package day1

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

func parseInput() ([]int, error) {
	file, err := os.Open("day1/input.txt")
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)

	var nums []int
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		nums = append(nums, num)
	}
	return nums, nil
}

func finalCount(nums []int) int {
	var count int
	prev := math.MaxInt
	for _, v := range nums {
		if v > prev {
			count++
		}
		prev = v
	}
	return count
}

// Part1 counts the number of times a depth measurement increases from
// the previous measurement
func Part1() (int, error) {
	nums, err := parseInput()
	if err != nil {
		return 0, err
	}
	return finalCount(nums), nil
}

// Part2 counts the number of times the sum of measurements in this sliding
// window increases from the previous sum
func Part2() (int, error) {
	nums, err := parseInput()
	if err != nil {
		return 0, err
	}

	var sums []int
	for i := 0; i < len(nums)-2; i++ {
		sum := nums[i] + nums[i+1] + nums[i+2]
		sums = append(sums, sum)
	}
	return finalCount(sums), nil
}
