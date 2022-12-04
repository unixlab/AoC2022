// Package day04 is our package for the 4th AoC day
package day04

import (
	"fmt"
	"strconv"
	"strings"
)

// RunPart1 is for the first star of the day
func RunPart1(day string, input []string) {
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
		}

		if elf2min <= elf1min && elf2max >= elf1max {
			fullyContainCounter++
		}

		// for whatever reason if the area is the same they do not
		// count as overlapping
		if elf1min == elf2min && elf1max == elf2max {
			fullyContainCounter--
		}
	}
	fmt.Printf("%s part 1 => %d\n", day, fullyContainCounter)
}

// RunPart2 is for the second star of the day
func RunPart2(day string, input []string) {
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
			continue
		}

		if elf2max >= elf1min && elf2max <= elf1max {
			totalOverlaps++
			continue
		}
	}
	fmt.Printf("%s part 2 => %d\n", day, totalOverlaps)
}
