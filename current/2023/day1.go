package main

import (
	"fmt"
	"strconv"
	"strings"
)

var numberNameMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

// First try for Day One, no Refactor.
func findFirstAndLastInt(strline string) []int {
	var int_pos []int

	for pos, char := range strline {
		if _, err := strconv.Atoi(string(char)); err == nil {
			int_pos = append(int_pos, pos)
		}
	}

	return int_pos

}

// Part Two

// Find first and last of the integers by name in a string
func findIntNames(strline string) map[string]int {
	var int_name_pos = make(map[string]int)

	for k, _ := range numberNameMap {

		first_index := strings.Index(strline, k)
		last_index := strings.LastIndex(strline, k)

		// Debug
		//fmt.Println(k, first_index, last_index)

		if first_index != -1 {
			save_name := k + "_firstindex"
			int_name_pos[save_name] = first_index
		}

		if last_index != -1 {
			save_name := k + "_lastindex"
			int_name_pos[save_name] = last_index
		}
	}

	// Debug
	// fmt.Println(int_name_pos)
	// os.Exit(0)

	return int_name_pos
}

func DayOne() {
	fmt.Println("Day One: \n----------")

	input := fileToString("input/day1.txt")

	////////////
	// PART ONE
	total_one := 0

	inputLines := strings.Split(input, "\n")

	for line := range inputLines {

		int_pos := findFirstAndLastInt(inputLines[line])
		int_str := ""

		if len(int_pos) > 0 {
			int_str += string(inputLines[line][int_pos[0]])
			int_str += string(inputLines[line][int_pos[len(int_pos)-1]])
		}

		final_int, err := strconv.Atoi(int_str)

		if err != nil {
			fmt.Println(err)
		}

		total_one += final_int

	}

	fmt.Println("Part One: ", total_one)

	////////////
	// PART TWO

	total_two := 0

	for line := range inputLines {

		low_pos := 0
		high_pos := 0
		low_type := "int"
		high_type := "int"

		low_int := ""
		high_int := ""

		int_names_pos := findIntNames(inputLines[line])
		int_pos := findFirstAndLastInt(inputLines[line])

		// If we have integers in the string, use them as the first and last
		if len(int_pos) > 0 {
			low_pos = int_pos[0]
			high_pos = int_pos[len(int_pos)-1]
			low_type = "int"
			high_type = "int"

		}

		// If we have integers by name, find the first and last
		// Rename type after the actual word we find
		for k, v := range int_names_pos {
			if v < low_pos {
				low_pos = v
				low_type = strings.Split(k, "_")[0]
			}

			if v > high_pos {
				high_pos = v
				high_type = strings.Split(k, "_")[0]
			}
		}

		// If we have integers by name, convert them to integers
		if low_type == "int" {
			new_int := string(inputLines[line][low_pos])

			low_int = new_int

		} else {
			low_int = numberNameMap[low_type]
		}

		if high_type == "int" {
			new_int := string(inputLines[line][high_pos])

			high_int = new_int

		} else {
			high_int = numberNameMap[high_type]
		}

		// Increment the total
		additional := low_int + high_int
		final_int, err := strconv.Atoi(additional)

		if err != nil {
			fmt.Println(err)
		}

		total_two += final_int

		// Debug
		//fmt.Println(low_int, low_type, high_int, high_type, additional, final_int, total_two)
	}

	fmt.Println("Part Two: ", total_two)
	fmt.Println()

}
