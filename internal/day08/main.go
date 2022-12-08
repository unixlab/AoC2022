// Package day08 is our package for the 8th AoC day
package day08

import (
	"strconv"
	"strings"
)

func visibleFromEdge(grid [][]int, x int, y int) bool {
	left := true
	for i := x - 1; i >= 0; i-- {
		if grid[y][x] <= grid[y][i] {
			left = false
		}
	}
	right := true
	for i := x + 1; i < len(grid[x]); i++ {
		if grid[y][x] <= grid[y][i] {
			right = false
		}
	}
	top := true
	for i := y - 1; i >= 0; i-- {
		if grid[y][x] <= grid[i][x] {
			top = false
		}
	}
	bottom := true
	for i := y + 1; i < len(grid); i++ {
		if grid[y][x] <= grid[i][x] {
			bottom = false
		}
	}
	if left || right || top || bottom {
		return true
	}
	return false
}

func countVisibleTrees(grid [][]int, x int, y int) int {
	left := 0
	for i := x - 1; i >= 0; i-- {
		left++
		if grid[y][x] <= grid[y][i] {
			break
		}
	}
	right := 0
	for i := x + 1; i < len(grid[x]); i++ {
		right++
		if grid[y][x] <= grid[y][i] {
			break
		}
	}
	top := 0
	for i := y - 1; i >= 0; i-- {
		top++
		if grid[y][x] <= grid[i][x] {
			break
		}
	}
	bottom := 0
	for i := y + 1; i < len(grid); i++ {
		bottom++
		if grid[y][x] <= grid[i][x] {
			break
		}
	}
	return left * right * top * bottom
}

// RunPart1 is for the first star of the day
func RunPart1(input []string) int {
	var grid [][]int
	for _, line := range input {
		var newRow []int
		for _, v := range strings.Split(line, "") {
			n, _ := strconv.Atoi(v)
			newRow = append(newRow, n)
		}
		grid = append(grid, newRow)
	}
	visibleTrees := 0
	for x := 1; x < len(grid)-1; x++ {
		for y := 1; y < len(grid[x])-1; y++ {
			if visibleFromEdge(grid, x, y) {
				visibleTrees++
			}
		}
	}
	return visibleTrees + 4*len(grid) - 4
}

// RunPart2 is for the second star of the day
func RunPart2(input []string) int {
	var grid [][]int
	for _, line := range input {
		var newRow []int
		for _, v := range strings.Split(line, "") {
			n, _ := strconv.Atoi(v)
			newRow = append(newRow, n)
		}
		grid = append(grid, newRow)
	}
	maxTreeScore := 0
	for y := range grid {
		for x := range grid[y] {
			c := countVisibleTrees(grid, x, y)
			if c > maxTreeScore {
				maxTreeScore = c
			}
		}
	}
	return maxTreeScore
}
