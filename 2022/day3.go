package main

import (
	"fmt"
	"strings"
)

// Splits the string for each rucksack in half and returns them in a slice
func getSackHalves(contents string) []string {
	var halves []string
	conlen := len(contents)

	halves = append(halves, contents[:conlen/2])
	halves = append(halves, contents[conlen/2:])

	return halves
}

// Compares two strings and determines which chars are the same
func compareSackHalves(sidea string, sideb string) string {
	matchstr := ""

	for a := range sidea {
		for b := range sideb {
			if sidea[a] == sideb[b] {
				matchstr = string(sidea[a])
			}
		}
	}

	return matchstr
}

// Returns the priority score for a given letter
func getLetterPriority(letter string) int {
	priority := 0
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for v := range letters {
		if string(letters[v]) == letter {
			priority = v + 1
		}
	}

	return priority
}

// Compares three stringsd and determiens which char is the same
func compareSackTrio(sidea string, sideb string, sidec string) string {
	var sideBMatches []string
	var sideCMatches []string
	matchstr := ""

	for a := range sidea {
		for b := range sideb {
			if sidea[a] == sideb[b] {
				sideBMatches = append(sideBMatches, string(sidea[a]))
			}
		}
		for c := range sidec {
			if sidea[a] == sidec[c] {
				sideCMatches = append(sideCMatches, string(sidea[a]))
			}
		}

	}

	matchstr = compareSackHalves(strings.Join(sideBMatches, ""), strings.Join(sideCMatches, ""))

	return matchstr
}

func dayThree() {
	input := inputToString("day3")
	splitput := strings.Split(input, "\n")

	//////////////
	// Part One //
	//////////////

	prioritySum := 0
	for v := range splitput {
		halves := getSackHalves(string(splitput[v]))
		wrongItem := compareSackHalves(halves[0], halves[1])
		priority := getLetterPriority(wrongItem)

		prioritySum = prioritySum + priority
	}

	fmt.Println("Day Three, Part One:", prioritySum)

	//////////////
	// Part Two //
	//////////////

	// Combine input into groups of three
	var groups [][]string
	var group []string

	for v := range splitput {
		group = append(group, splitput[v])

		if len(group) >= 3 {
			groups = append(groups, group)
			group = nil
		}
	}

	badgeSum := 0
	for v := range groups {
		badge := compareSackTrio(groups[v][0], groups[v][1], groups[v][2])

		badgeSum = badgeSum + getLetterPriority(badge)
	}

	fmt.Println("Day Three, Part One:", badgeSum)

}
