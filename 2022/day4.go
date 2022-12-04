package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Compares the two number ranges to see if one fully includes the other
// Return True or False
func comparePairs(paira [2]int, pairb [2]int) bool {

	pa1 := paira[0]
	pa2 := paira[1]
	pb1 := pairb[0]
	pb2 := pairb[1]

	if pa1 >= pb1 {
		if pa2 <= pb2 {
			return true
		}
	}

	if pb1 >= pa1 {
		if pb2 <= pa2 {
			return true
		}
	}

	return false
}

// Compares the two number ranges to see if one includes any part of the other
// Return True or False
func comparePairsFull(paira [2]int, pairb [2]int) bool {

	pa1 := paira[0]
	pa2 := paira[1]
	pb1 := pairb[0]
	pb2 := pairb[1]

	if pb1 <= pa2 {
		if pa1 <= pb1 {
			return true
		}
	}

	if pa1 <= pb2 {
		if pb1 <= pa1 {
			return true
		}
	}

	return false
}

func dayFour() {
	input := inputToString("day4")
	splitput := strings.Split(input, "\n")

	pairsSum := 0
	pairSumFull := 0

	for v := range splitput {
		pairs := strings.Split(splitput[v], ",")
		p1 := strings.Split(pairs[0], "-")
		p2 := strings.Split(pairs[1], "-")

		p1a, _ := strconv.Atoi(p1[0])
		p1b, _ := strconv.Atoi(p1[1])
		p2a, _ := strconv.Atoi(p2[0])
		p2b, _ := strconv.Atoi(p2[1])

		paira := [2]int{p1a, p1b}
		pairb := [2]int{p2a, p2b}

		if comparePairs(paira, pairb) {
			pairsSum = pairsSum + 1
		}

		if comparePairsFull(paira, pairb) {
			pairSumFull = pairSumFull + 1
		}

	}

	fmt.Println("Day Four, Part One:", pairsSum)
	fmt.Println("Day Four, Part Two:", pairSumFull)

}
