package main

import (
	"fmt"
	"strconv"
	"strings"
)

////////////////////////////////////////////////////////////////

// Checks if a tree is visible across a row
func checkHorizontal(row string, treeIndex int) bool {
	rowLen := len(row)
	b, err := strconv.Atoi(string(row[treeIndex]))

	// Edges are visible by default
	if treeIndex == rowLen-1 {
		return true
	}

	if treeIndex == 0 {
		return true
	}

	for i := 0; i < rowLen; i++ {
		if i != treeIndex {
			a, err := strconv.Atoi(string(row[i]))

			if err != nil {
				fmt.Println(err)
			}

			if a >= b {
				return false
			}

		}
	}

	if err != nil {
		fmt.Println(err)
	}

	return true
}

////////////////////////////////////////////////////////////////

func checkVertical(input string, currentRow int, treeIndex int) bool {
	rows := strings.Split(input, "\n")

	for k, v := range rows {
		rowLen := len(v)
		a, err := strconv.Atoi(string(v[treeIndex]))
		b, err := strconv.Atoi(string(string(input[currentRow])[treeIndex]))

		if err != nil {
			fmt.Println(err)
		}

		if k != currentRow {
			// Edges are visible by default
			if treeIndex == rowLen-1 {
				return true
			}

			if treeIndex == 0 {
				return true
			}

			if a >= b {
				return false
			}

		}
	}

	return true
}

////////////////////////////////////////////////////////////////

func dayEight() {
	input := inputToString("day8")

	totalVisible := 0

	for k, v := range strings.Split(input, "\n") {
		for kt := range v {
			if checkHorizontal(v, kt) {
				totalVisible++
				break
			}

			if checkVertical(v, k, kt) {
				totalVisible++
				break
			}
		}

	}

	fmt.Println("Day Eight, Part One:", totalVisible)

}
