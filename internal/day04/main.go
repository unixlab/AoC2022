// Package day04 is our package for the 4th AoC day
package day04

import (
	"strconv"
	"strings"
)

// RunPart1 is for the first star of the day
func RunPart1(input []string) int {
	fullyContainCounter := 0
	for _, line := range input {
		assignments := strings.Split(line, ",")
		elf1 := strings.Split(assignments[0], "-")
		elf2 := strings.Split(assignments[1], "-")

		elf1min, _ := strconv.Atoi(elf1[0])
		elf1max, _ := strconv.Atoi(elf1[1])
		elf2min, _ := strconv.Atoi(elf2[0])
		elf2max, _ := strconv.Atoi(elf2[1])

		if elf1min <= elf2min && elf1max >= elf2max {
			fullyContainCounter++
			continue
		}

		if elf2min <= elf1min && elf2max >= elf1max {
			fullyContainCounter++
		}
	}
	return fullyContainCounter
}

// RunPart2 is for the second star of the day
func RunPart2(input []string) int {
	totalOverlaps := 0
	for _, line := range input {
		assignments := strings.Split(line, ",")
		elf1 := strings.Split(assignments[0], "-")
		elf2 := strings.Split(assignments[1], "-")

		elf1min, _ := strconv.Atoi(elf1[0])
		elf1max, _ := strconv.Atoi(elf1[1])
		elf2min, _ := strconv.Atoi(elf2[0])
		elf2max, _ := strconv.Atoi(elf2[1])

		if elf1min >= elf2min && elf1min <= elf2max {
			totalOverlaps++
			continue
		}

		if elf1max >= elf2min && elf1max <= elf2max {
			totalOverlaps++
			continue
		}

		if elf2min >= elf1min && elf2min <= elf1max {
			totalOverlaps++
		}
	}
	return totalOverlaps
}
