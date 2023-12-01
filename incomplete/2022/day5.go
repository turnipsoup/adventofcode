package main

import (
	"fmt"
	"strconv"
	"strings"
)

type loadingDock struct {
	One   *[]string
	Two   *[]string
	Three *[]string
	Four  *[]string
	Five  *[]string
	Six   *[]string
	Seven *[]string
	Eight *[]string
	Nine  *[]string
}

// Moves one package from the top of Pile One to Pile Two
// Pile One is the FROM pile, with Pile Two being the DEPOSITED pile
func movePackages(pileOne int, pileTwo int, loadingDockMap map[int][]string) {

	// Get the last item in the FROM pile
	moveItem := loadingDockMap[pileOne][len(loadingDockMap[pileOne])-1]

	// Remove it
	if len(loadingDockMap[pileOne]) > 0 {
		loadingDockMap[pileOne] = loadingDockMap[pileOne][:len(loadingDockMap[pileOne])-1]
	}

	// Append it to the target pile
	loadingDockMap[pileTwo] = append(loadingDockMap[pileTwo], moveItem)

}

// Moves multiple packages from the top of Pile One to Pile Two
// Pile One is the FROM pile, with Pile Two being the DEPOSITED pile
func moveMultiplePackages(pileOne int, pileTwo int, numMove int, loadingDockMap map[int][]string) {

	if len(loadingDockMap[pileOne])-numMove > 0 {
		// Get the requested item in the FROM pile
		moveItems := loadingDockMap[pileOne][:len(loadingDockMap[pileOne])-1-numMove]

		// Remove it
		if len(loadingDockMap[pileOne]) > 0 {
			loadingDockMap[pileOne] = loadingDockMap[pileOne][:len(loadingDockMap[pileOne])-1-numMove]
		}

		// Append it to the target pile
		for _, v := range moveItems {
			loadingDockMap[pileTwo] = append(loadingDockMap[pileTwo], v)
		}
	}
}

func dayFivePartOne() {
	// Load input
	input := inputToString("day5")
	packageMovements := strings.Split(input, "\n\n")[1]

	// Set up the loading dock stuct
	dock := &loadingDock{
		One:   &[]string{"Z", "J", "G"},
		Two:   &[]string{"Q", "L", "R", "P", "W", "F", "V", "C"},
		Three: &[]string{"F", "P", "M", "C", "L", "G", "R"},
		Four:  &[]string{"L", "F", "B", "W", "P", "H", "M"},
		Five:  &[]string{"G", "C", "F", "S", "V", "Q"},
		Six:   &[]string{"W", "H", "J", "Z", "M", "Q", "T", "L"},
		Seven: &[]string{"H", "F", "S", "B", "W"},
		Eight: &[]string{"F", "J", "Z", "S"},
		Nine:  &[]string{"M", "C", "D", "P", "F", "H", "B", "T"},
	}

	var loadingDockMap = map[int][]string{
		1: *dock.One,
		2: *dock.Two,
		3: *dock.Three,
		4: *dock.Four,
		5: *dock.Five,
		6: *dock.Six,
		7: *dock.Seven,
		8: *dock.Eight,
		9: *dock.Nine,
	}

	for _, v := range strings.Split(packageMovements, "\n") {
		splitRow := strings.Split(v, " ")
		numMove, err := strconv.Atoi(splitRow[1])
		fromPile, err := strconv.Atoi(splitRow[3])
		toPile, err := strconv.Atoi(splitRow[5])

		if err != nil {
			fmt.Println(err)
		}

		for i := 0; i < numMove; i++ {
			movePackages(fromPile, toPile, loadingDockMap)
		}

	}

	fmt.Println("Day Five, part One:", loadingDockMap)
}

func dayFivePartTwo() {
	// Load input
	input := inputToString("day5")
	packageMovements := strings.Split(input, "\n\n")[1]

	// Set up the loading dock stuct
	dock := &loadingDock{
		One:   &[]string{"Z", "J", "G"},
		Two:   &[]string{"Q", "L", "R", "P", "W", "F", "V", "C"},
		Three: &[]string{"F", "P", "M", "C", "L", "G", "R"},
		Four:  &[]string{"L", "F", "B", "W", "P", "H", "M"},
		Five:  &[]string{"G", "C", "F", "S", "V", "Q"},
		Six:   &[]string{"W", "H", "J", "Z", "M", "Q", "T", "L"},
		Seven: &[]string{"H", "F", "S", "B", "W"},
		Eight: &[]string{"F", "J", "Z", "S"},
		Nine:  &[]string{"M", "C", "D", "P", "F", "H", "B", "T"},
	}

	var loadingDockMap = map[int][]string{
		1: *dock.One,
		2: *dock.Two,
		3: *dock.Three,
		4: *dock.Four,
		5: *dock.Five,
		6: *dock.Six,
		7: *dock.Seven,
		8: *dock.Eight,
		9: *dock.Nine,
	}

	for _, v := range strings.Split(packageMovements, "\n") {
		splitRow := strings.Split(v, " ")
		numMove, err := strconv.Atoi(splitRow[1])
		fromPile, err := strconv.Atoi(splitRow[3])
		toPile, err := strconv.Atoi(splitRow[5])

		if err != nil {
			fmt.Println(err)
		}

		moveMultiplePackages(fromPile, toPile, numMove, loadingDockMap)

	}

	fmt.Println("Day Five, part Two:", loadingDockMap)
}

func dayFive() {

	dayFivePartOne()
	dayFivePartTwo()

}
