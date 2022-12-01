package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Returns the sum of the values of the clories an elf is carrying
func sumCalories(inlist string) int {
	inp := strings.Split(inlist, "\n")
	calsum := 0

	for _, v := range inp {

		inpi, _ := strconv.Atoi(v)

		calsum = calsum + inpi
	}

	return calsum

}

// Gets largest integer from an array of integers
func getLargestInt(ints []int) int {
	large := 0

	for _, v := range ints {
		if v > large {
			large = v
		}
	}

	return large

}

func dayOne() {

	// Part One
	input := inputToString("day1")
	splitput := strings.Split(input, "\n\n")

	var calsums []int

	for _, v := range splitput {
		calsums = append(calsums, sumCalories(v))
	}

	mostCals := getLargestInt(calsums)

	fmt.Println("Day One, Part One:", mostCals)

	/////

	// Part Two

	// Sort it
	sort.Ints(calsums)

	// Reverse it
	for i := len(calsums)/2 - 1; i >= 0; i-- {
		opp := len(calsums) - 1 - i
		calsums[i], calsums[opp] = calsums[opp], calsums[i]
	}

	fmt.Println("Day One, Part Two:", calsums[0]+calsums[1]+calsums[2])

}
