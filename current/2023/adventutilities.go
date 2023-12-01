package main

import (
	"fmt"
	"os"
)

///////

// Read a file and return it as a string
func fileToString(filename string) string {
	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	return string(file)

}

// Return the larger of two integers
func largerInt(i int, j int) int {
	if i >= j {
		return i
	}

	return j
}
