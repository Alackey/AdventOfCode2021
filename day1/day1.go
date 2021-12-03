package day1

import (
	"math"
)

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
func Part1(nums []int) (int, error) {
	return finalCount(nums), nil
}

// Part2 counts the number of times the sum of measurements in this sliding
// window increases from the previous sum
func Part2(nums []int) (int, error) {
	var sums []int
	for i := 0; i < len(nums)-2; i++ {
		sum := nums[i] + nums[i+1] + nums[i+2]
		sums = append(sums, sum)
	}
	return finalCount(sums), nil
}
