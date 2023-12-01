package main

import (
	"fmt"
	"strconv"
	"strings"
)

////////////////////////////////////////////////////////////////

// Checks if a tree is visible across a row
func checkHorizontal(row []string, treeIndex int) bool {
	rowLen := len(row)
	b, err := strconv.Atoi(string(row[treeIndex]))

	// Edges are visible by default
	if treeIndex >= rowLen-1 {
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

			if a > b {
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

func checkVertical(splitput []string, currentRow int, treeIndex int) bool {
	for _, v := range splitput {
		rowLen := len(v)
		rowArr := strings.Split(v, "")

		// Edges are visible by default
		if treeIndex >= rowLen-1 {
			return true
		}

		if treeIndex == 0 {
			return true
		}

		if len(rowArr) < 1 {
			continue
		}
		a, err := strconv.Atoi(rowArr[treeIndex])
		b, err := strconv.Atoi(string(string(splitput[currentRow])[treeIndex]))

		if err != nil {
			fmt.Println(err)
		}

		if a > b {
			return false
		}

	}
	return true
}

////////////////////////////////////////////////////////////////

func dayEight() {
	input := inputToString("day8")
	splitput := strings.Split(input, "\n")
	totalVisible := 0

	for i := 0; i < len(splitput); i++ {
		rowArr := strings.Split(splitput[i], "")
		for t := 0; t < len(rowArr); t++ {
			fmt.Println("Row:", i, "Char:", t, rowArr[t], rowArr)
			if checkHorizontal(rowArr, t) {
				totalVisible++
			}
			if checkVertical(splitput, i, t) {
				totalVisible++
			}
		}
	}

	fmt.Println("Day Eight, Part One:", totalVisible)

}
