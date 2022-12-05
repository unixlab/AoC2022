// Package day05 is our package for the 5th AoC day
package day05

import (
	"fmt"
	"strconv"
	"strings"
)

func parseMove(instruction string) (int, int, int) {
	splits := strings.Split(instruction, " ")
	count, _ := strconv.Atoi(splits[1])
	src, _ := strconv.Atoi(splits[3])
	src--
	dst, _ := strconv.Atoi(splits[5])
	dst--
	return count, src, dst
}

// RunPart1 is for the first star of the day
func RunPart1(day string, input []string) {
	var stack [][]string
	for _, line := range input {
		if strings.HasPrefix(line, "move") {
			count, src, dst := parseMove(line)
			for i := 0; i < count; i++ {
				stack[dst] = append([]string{stack[src][0]}, stack[dst]...)
				stack[src] = stack[src][1:]
			}
		} else {
			stackIndex := 0
			for i := 1; i < len(line); i += 4 {
				crate := line[i : i+1]
				if crate[0] < 60 {
					stackIndex++
					continue
				}
				for i := len(stack) - 1; i < stackIndex; i++ {
					stack = append(stack, []string{})
				}
				stack[stackIndex] = append(stack[stackIndex], crate)
				stackIndex++
			}
		}
	}
	var result strings.Builder
	for i := 0; i < len(stack); i++ {
		result.WriteString(stack[i][0])
	}
	fmt.Printf("%s part 1 => %s\n", day, result.String())
}

// RunPart2 is for the second star of the day
func RunPart2(day string, input []string) {
	var stack [][]string
	for _, line := range input {
		if strings.HasPrefix(line, "move") {
			count, src, dst := parseMove(line)
			var newDst []string
			newDst = append(newDst, stack[src][:count]...)
			newDst = append(newDst, stack[dst]...)
			stack[dst] = newDst
			stack[src] = stack[src][count:]
		} else {
			stackIndex := 0
			for i := 1; i < len(line); i += 4 {
				crate := line[i : i+1]
				if crate[0] < 60 {
					stackIndex++
					continue
				}
				for i := len(stack) - 1; i < stackIndex; i++ {
					stack = append(stack, []string{})
				}
				stack[stackIndex] = append(stack[stackIndex], crate)
				stackIndex++
			}
		}
	}
	var result strings.Builder
	for i := 0; i < len(stack); i++ {
		result.WriteString(stack[i][0])
	}
	fmt.Printf("%s part 2 => %s\n", day, result.String())
}
