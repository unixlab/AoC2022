// Package day01 is our package for the 1st AoC day
package day01

import (
	"sort"
	"strconv"
)

// RunPart1 is for the first star of the day
func RunPart1(input []string) int {
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
	return maxCalories
}

// RunPart2 is for the second star of the day
func RunPart2(input []string) int {
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
	elfCalories = append(elfCalories, current)
	sort.Ints(elfCalories)
	maxThree := 0
	for i := len(elfCalories) - 1; i >= len(elfCalories)-3; i-- {
		maxThree += elfCalories[i]
	}
	return maxThree
}
