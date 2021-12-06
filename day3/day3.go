package day3

import (
	"strconv"
)

// Part1 returns the power consumption (gamma * epsilon rate) of the submarine
func Part1(lines []string) (int64, error) {
	sums := calculateSums(lines)

	for _, line := range lines {
		for i, char := range line {
			switch char {
			case '1':
				sums[i]++
			case '0':
				sums[i]--
			}
		}
	}

	var gammaBits, epsilonBits string
	for _, sum := range sums {
		if sum > 0 {
			gammaBits += "1"
			epsilonBits += "0"
		} else {
			gammaBits += "0"
			epsilonBits += "1"
		}
	}

	gamma, err := strconv.ParseInt(gammaBits, 2, 64)
	if err != nil {
		return 0, err
	}

	epsilon, err := strconv.ParseInt(epsilonBits, 2, 64)
	if err != nil {
		return 0, err
	}
	return gamma * epsilon, nil
}

// Part2 returns the life support rating (oxygen generator * CO2 scrubber rating) of the submarine
func Part2(lines []string) (int64, error) {
	sums := calculateSums(lines)

	// Oxygen
	oxygenRatings := make([]string, len(lines))
	copy(oxygenRatings, lines)

	var index int
	for len(oxygenRatings) != 1 {
		if sums[index] >= 0 {
			var idx int
			for _, rating := range oxygenRatings {
				if rating[index] == '0' {
					oxygenRatings = removeIndex(oxygenRatings, idx)
				} else {
					idx++
				}
			}
		} else {
			var idx int
			for _, rating := range oxygenRatings {
				if rating[index] == '1' {
					oxygenRatings = removeIndex(oxygenRatings, idx)
				} else {
					idx++
				}
			}
		}
		index++
		sums = calculateSums(oxygenRatings)
	}

	// CO2 Scrubber
	scrubberRatings := make([]string, len(lines))
	copy(scrubberRatings, lines)

	index = 0
	for len(scrubberRatings) != 1 {
		if sums[index] >= 0 {
			var idx int
			for _, rating := range scrubberRatings {
				if rating[index] == '1' {
					scrubberRatings = removeIndex(scrubberRatings, idx)
				} else {
					idx++
				}
			}
		} else {
			var idx int
			for _, rating := range scrubberRatings {
				if rating[index] == '0' {
					scrubberRatings = removeIndex(scrubberRatings, idx)
				} else {
					idx++
				}
			}
		}
		index++
		sums = calculateSums(scrubberRatings)
	}

	oxygen, err := strconv.ParseInt(oxygenRatings[0], 2, 64)
	if err != nil {
		return 0, err
	}

	scrubber, err := strconv.ParseInt(scrubberRatings[0], 2, 64)
	if err != nil {
		return 0, err
	}
	return oxygen * scrubber, nil
}

func calculateSums(nums []string) []int {
	var sums []int
	for i := 0; i < len(nums[0]); i++ {
		sums = append(sums, 0)
	}

	for _, line := range nums {
		for i, char := range line {
			switch char {
			case '1':
				sums[i]++
			case '0':
				sums[i]--
			}
		}
	}
	return sums
}

func removeIndex(s []string, index int) []string {
	ret := make([]string, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
