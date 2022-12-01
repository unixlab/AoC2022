package day01

import (
	"fmt"
	"sort"
	"strconv"
)

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
