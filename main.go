package main

import (
	"fmt"

	"github.com/alackey/adventofcode2021/day1"
)

func main() {
	count, err := day1.Part1()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Total Part 1: %d\n", count)

	count, err = day1.Part2()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Total Part 2: %d\n", count)
}
