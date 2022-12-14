// Package day14 is our package for the 14th AoC day
package day14

import (
	"math"
	"strconv"
	"strings"
)

// absDiff calculates the absolut difference as int
func absDiff(x int, y int) int {
	return int(math.Abs(float64(x) - float64(y)))
}

func buildGrid(input []string) ([200][1000]int, int) {
	var grid [200][1000]int
	maxY := -1
	for _, line := range input {
		line = strings.Replace(line, " ", "", -1)
		oldX, oldY := -1, -1
		for _, step := range strings.Split(line, "->") {
			cords := strings.Split(step, ",")
			x, _ := strconv.Atoi(cords[0])
			y, _ := strconv.Atoi(cords[1])
			if oldX == -1 && oldY == -1 {
				oldX = x
				oldY = y
				grid[oldY][oldX] = 1
				continue
			}
			for absDiff(oldX, x) > 0 {
				if oldX < x {
					oldX++
				} else {
					oldX--
				}
				grid[oldY][oldX] = 1
			}
			for absDiff(oldY, y) > 0 {
				if oldY < y {
					oldY++
				} else {
					oldY--
				}
				grid[oldY][oldX] = 1
			}
			if oldY > maxY {
				maxY = oldY
			}
		}
	}
	maxY += 2
	return grid, maxY
}

// RunPart1 is for the first star of the day
func RunPart1(input []string) int {
	grid, _ := buildGrid(input)
	sandCounter := 0
	for {
		breakLoop := false
		sandY := 0
		sandX := 500
		for {
			if sandY+1 >= len(grid) {
				breakLoop = true
				break
			}
			if grid[sandY+1][sandX] == 0 {
				sandY++
			} else if grid[sandY+1][sandX-1] == 0 {
				sandY++
				sandX--
			} else if grid[sandY+1][sandX+1] == 0 {
				sandY++
				sandX++
			} else {
				grid[sandY][sandX] = 2
				break
			}
		}
		if breakLoop {
			break
		}
		sandCounter++
	}
	return sandCounter
}

// RunPart2 is for the first star of the day
func RunPart2(input []string) int {
	grid, maxY := buildGrid(input)
	for x := 0; x < len(grid[maxY]); x++ {
		grid[maxY][x] = 1
	}
	sandCounter := 0
	for {
		sandY := 0
		sandX := 500
		for {
			if grid[sandY+1][sandX] == 0 {
				sandY++
			} else if grid[sandY+1][sandX-1] == 0 {
				sandY++
				sandX--
			} else if grid[sandY+1][sandX+1] == 0 {
				sandY++
				sandX++
			} else {
				grid[sandY][sandX] = 2
				break
			}
		}
		sandCounter++
		if grid[0][500] != 0 {
			break
		}
	}
	return sandCounter
}
