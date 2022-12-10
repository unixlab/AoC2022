// Package day10 is our package for the 10th AoC day
package day10

import (
	"math"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

// Instruction struct for parsed instructions
type Instruction struct {
	Cycles int
	Value  int
}

// absDiff calculates the absolut difference as int
func absDiff(x int, y int) int {
	return int(math.Abs(float64(x) - float64(y)))
}

// Run as two return values and can be used for both parts
func Run(input []string, printColor bool) (int, string) {
	var instructions []Instruction
	for _, line := range input {
		var newInstruction Instruction
		switch strings.Split(line, " ")[0] {
		case "addx":
			newInstruction.Cycles = 2
			newInstruction.Value, _ = strconv.Atoi(strings.Split(line, " ")[1])
		case "noop":
			newInstruction.Cycles = 1
			newInstruction.Value = 0
		default:
			return -1, ""
		}
		instructions = append(instructions, newInstruction)
	}
	idx := 0
	regX := 1
	signalStrength := 0
	spriteLit := "#"
	spriteDark := "."
	if printColor {
		spriteLit = color.New(color.BgWhite).Sprint(" ")
		spriteDark = color.New(color.BgBlack).Sprint(" ")
	}
	var sprite strings.Builder
	for cycle := 1; cycle <= 240; cycle++ {
		if absDiff((cycle-1)%40, regX) < 2 {
			sprite.WriteString(spriteLit)
		} else {
			sprite.WriteString(spriteDark)
		}
		if cycle%40 == 0 {
			sprite.WriteString("\n")
		}
		if cycle%40 == 20 {
			signalStrength += regX * cycle
		}
		instructions[idx].Cycles--
		if instructions[idx].Cycles == 0 {
			regX += instructions[idx].Value
			idx++
		}
	}
	return signalStrength, sprite.String()
}
