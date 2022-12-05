// Package day02 is our package for the 2nd AoC day
package day02

import (
	"strings"
)

// The score for a single round is the score for the shape you selected plus the score for the outcome of the round
const (
	Rock     = 1
	Paper    = 2
	Scissors = 3
	Loose    = 0
	Draw     = 3
	Win      = 6
)

// RunPart1 is for the first star of the day
func RunPart1(input []string) int {
	score := 0
	for _, line := range input {
		actions := strings.Split(line, " ")
		switch actions[0] {
		case "A":
			switch actions[1] {
			case "Y":
				score += Paper
				score += Win
			case "X":
				score += Rock
				score += Draw
			case "Z":
				score += Scissors
				score += Loose
			}
		case "B":
			switch actions[1] {
			case "Y":
				score += Paper
				score += Draw
			case "X":
				score += Rock
				score += Loose
			case "Z":
				score += Scissors
				score += Win
			}
		case "C":
			switch actions[1] {
			case "Y":
				score += Paper
				score += Loose
			case "X":
				score += Rock
				score += Win
			case "Z":
				score += Scissors
				score += Draw
			}
		}
	}
	return score
}

// RunPart2 is for the second star of the day
func RunPart2(input []string) int {
	score := 0
	for _, line := range input {
		actions := strings.Split(line, " ")
		switch actions[0] {
		case "A":
			switch actions[1] {
			case "Y":
				score += Draw
				score += Rock
			case "X":
				score += Loose
				score += Scissors
			case "Z":
				score += Win
				score += Paper
			}
		case "B":
			switch actions[1] {
			case "Y":
				score += Draw
				score += Paper
			case "X":
				score += Loose
				score += Rock
			case "Z":
				score += Win
				score += Scissors
			}
		case "C":
			switch actions[1] {
			case "Y":
				score += Draw
				score += Scissors
			case "X":
				score += Loose
				score += Paper
			case "Z":
				score += Win
				score += Rock
			}
		}
	}
	return score
}
