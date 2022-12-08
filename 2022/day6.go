package main

import (
	"fmt"
)

// Looks at four chars at a time through an entire string
// Returns the ending index of the first set of four
// That has four unique characters
func findMarker(input string, markerLen int) int {
	inputLen := len(input)
	markerIndex := 0

	if inputLen < markerLen {
		return 0
	}

	for i := 0; i <= inputLen-markerLen-1; i++ {
		//fmt.Println(input[i : i+4])
		if uniqueString(input[i : i+markerLen]) {
			return i + markerLen
		}
	}

	return markerIndex
}

// Checks if a single string is fully unique
func uniqueString(arr string) bool {
	m := make(map[rune]bool)
	for _, i := range arr {
		_, ok := m[i]
		if ok {
			return false
		}

		m[i] = true
	}

	return true
}

func daySix() {
	input := inputToString("day6")
	fmt.Println("Day Six, Part One:", findMarker(input, 4))
	fmt.Println("Day Six, Part Two:", findMarker(input, 14))
}
