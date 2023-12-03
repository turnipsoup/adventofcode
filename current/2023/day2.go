package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Variables and Globals
var total_cubes = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

type BagDraw struct {
	Red   int
	Green int
	Blue  int
}

type RoundMaxCount struct {
	Red   int
	Green int
	Blue  int
}

func DayTwo() {

	input := fileToString("input/day2.txt")
	games := strings.Split(input, "\n")

	var valid_games_sum = 0
	var min_power_sum = 0

	for g := range games {

		var maxCount = RoundMaxCount{
			Red:   0,
			Green: 0,
			Blue:  0,
		}

		game_valid := true
		gsplit := strings.Split(games[g], ": ")
		gameid, err := strconv.Atoi(strings.Split(gsplit[0], " ")[1])

		if err != nil {
			fmt.Println("Error parsing gameid", err)
		}

		rounds := strings.Split(gsplit[1], "; ")

		for round := range rounds {
			var bagdraw BagDraw

			cubes := strings.Split(rounds[round], ",")

			for cube := range cubes {
				trim_cube := strings.Trim(cubes[cube], " ")
				cube_split := strings.Split(trim_cube, " ")
				cube_color := cube_split[1]
				cube_amount, err := strconv.Atoi(cube_split[0])

				// DEBUG
				//fmt.Println(cube_color, cube_amount, cube_split)

				if err != nil {
					fmt.Println("Error parsing cube amount", err)
				}

				switch cube_color {
				case "red":
					bagdraw.Red += cube_amount
					if cube_amount > maxCount.Red {
						maxCount.Red = cube_amount
					}
				case "green":
					bagdraw.Green += cube_amount
					if cube_amount > maxCount.Green {
						maxCount.Green = cube_amount
					}
				case "blue":
					bagdraw.Blue += cube_amount
					if cube_amount > maxCount.Blue {
						maxCount.Blue = cube_amount
					}
				default:
					fmt.Println("Invalid color: ", cube_color)
				}

				if bagdraw.Red > total_cubes["red"] || bagdraw.Green > total_cubes["green"] || bagdraw.Blue > total_cubes["blue"] {
					game_valid = false
				}

			}
		}

		if game_valid {
			valid_games_sum += gameid
		}

		if maxCount.Red == 0 {
			maxCount.Red = 1
		}
		if maxCount.Green == 0 {
			maxCount.Green = 1
		}
		if maxCount.Blue == 0 {
			maxCount.Blue = 1
		}

		countpower := maxCount.Red * maxCount.Green * maxCount.Blue
		min_power_sum += countpower
		//fmt.Println(gameid, maxCount, countpower, min_power_sum)
	}

	fmt.Println("Day Two, Part One: ", valid_games_sum)
	fmt.Println("Day Two, Part Two: ", min_power_sum)

}
