// Package day02 is our package for second AoC day
package day02

import (
	"fmt"
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
func RunPart1(day string, input []string) {
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
			default:
				panic("unknown action")
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
			default:
				panic("unknown action")
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
			default:
				panic("unknown action")
			}
		default:
			panic("unknown action")
		}
	}
	fmt.Printf("%s part 1 => %d\n", day, score)
}

// RunPart2 is for the second star of the day
func RunPart2(day string, input []string) {
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
			default:
				panic("unknown action")
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
			default:
				panic("unknown action")
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
			default:
				panic("unknown action")
			}
		default:
			panic("unknown action")
		}
	}
	fmt.Printf("%s part 2 => %d\n", day, score)
}
