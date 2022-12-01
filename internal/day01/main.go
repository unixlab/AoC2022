// Package day01 is our package for first AoC day
package day01

import (
	"fmt"
	"sort"
	"strconv"
)

// RunPart1 is for the first star of the day
func RunPart1(day string, input []string) {
	currentCalories := 0
	maxCalories := 0
	for _, line := range input {
		if line == "" {
			if currentCalories > maxCalories {
				maxCalories = currentCalories
			}
			currentCalories = 0
		} else {
			lineCalories, _ := strconv.Atoi(line)
			currentCalories += lineCalories
		}
	}
	fmt.Printf("%s part 1 => %d\n", day, maxCalories)
}

// RunPart2 is for the second star of the day
func RunPart2(day string, input []string) {
	current := 0
	var elfCalories []int
	for _, line := range input {
		if line == "" {
			elfCalories = append(elfCalories, current)
			current = 0
		} else {
			lineCalories, _ := strconv.Atoi(line)
			current += lineCalories
		}
	}
	sort.Ints(elfCalories)
	maxThree := 0
	for i := len(elfCalories) - 1; i >= len(elfCalories)-3; i-- {
		maxThree += elfCalories[i]
	}
	fmt.Printf("%s part 2 => %d\n", day, maxThree)
}
