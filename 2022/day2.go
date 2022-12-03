package main

import (
	"fmt"
	"strings"
)

type ItemPoints struct {
	Rock     int
	Paper    int
	Scissors int
}

type GameOutcomePoints struct {
	Win  int
	Loss int
	Draw int
}

func determineRPSWinner(player string, opponent string, scores GameOutcomePoints) int {

	if player == "X" {
		if opponent == "C" {
			return scores.Win
		}
		if opponent == "B" {
			return scores.Loss
		}
		if opponent == "A" {
			return scores.Draw
		}
	}

	if player == "Y" {
		if opponent == "A" {
			return scores.Win
		}
		if opponent == "C" {
			return scores.Loss
		}
		if opponent == "B" {
			return scores.Draw
		}
	}

	if player == "Z" {
		if opponent == "B" {
			return scores.Win
		}
		if opponent == "A" {
			return scores.Loss
		}
		if opponent == "C" {
			return scores.Draw
		}
	}

	// If all else fails return a draw I guess
	return scores.Draw

}

func getRPSTotalGameScore(game string, responseitems map[string]int, scores GameOutcomePoints) int {
	players := strings.Split(game, " ")
	opponent := players[0]
	player := players[1]

	playerPoints := responseitems[player]
	outcomePoints := determineRPSWinner(player, opponent, scores)

	return playerPoints + outcomePoints
}

func determineRPSThrow(game string) string {

	players := strings.Split(game, " ")
	opponentMove := players[0]
	outcome := players[1]

	yourMove := "X"

	if opponentMove == "A" {
		if outcome == "X" {
			yourMove = "Z"
		}
		if outcome == "Y" {
			yourMove = "X"
		}
		if outcome == "Z" {
			yourMove = "Y"
		}
	}

	if opponentMove == "B" {
		if outcome == "X" {
			yourMove = "X"
		}
		if outcome == "Y" {
			yourMove = "Y"
		}
		if outcome == "Z" {
			yourMove = "Z"
		}
	}

	if opponentMove == "C" {
		if outcome == "X" {
			yourMove = "Y"
		}
		if outcome == "Y" {
			yourMove = "Z"
		}
		if outcome == "Z" {
			yourMove = "X"
		}
	}

	// If all else fails return a draw I guess
	return fmt.Sprintf("%s %s", opponentMove, yourMove)
}

func dayTwo() {

	//// Part One

	// Easily keep track of what we have
	ItemScore := ItemPoints{1, 2, 3}
	GamePoints := GameOutcomePoints{6, 0, 3}

	// Map symbols to points
	// ItemSymbols := map[string]int{
	// 	"A": ItemScore.Rock,
	// 	"B": ItemScore.Paper,
	// 	"C": ItemScore.Scissors,
	// }

	ResponseSymbols := map[string]int{
		"X": ItemScore.Rock,
		"Y": ItemScore.Paper,
		"Z": ItemScore.Scissors,
	}

	inputs := inputToString("day2")
	//inputs := "A Y\nB X\nC Z" // Test String
	splitputs := strings.Split(inputs, "\n")

	totalScore := 0

	for v := range splitputs {
		sc := getRPSTotalGameScore(splitputs[v], ResponseSymbols, GamePoints)
		totalScore = totalScore + sc
	}

	fmt.Println("Day Two, Part One:", totalScore)

	//// Part Two

	throwScore := 0

	for v := range splitputs {
		game := determineRPSThrow(splitputs[v])
		sc := getRPSTotalGameScore(game, ResponseSymbols, GamePoints)
		//fmt.Println(splitputs[v], game, "--", sc)
		throwScore = throwScore + sc
	}

	fmt.Println("Day Two, Part One:", throwScore)

}
