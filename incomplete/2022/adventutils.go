package main

import (
	"fmt"
	"os"
)

// Files
//

// Open the input file as a string
func inputToString(day string) string {
	inptext := ""
	inpfile := fmt.Sprintf("%s/%s.txt", INPUTDIR, day)

	inp, err := os.ReadFile(inpfile)

	if err != nil {
		fmt.Println(err)
	}

	inptext = string(inp)

	return inptext
}
