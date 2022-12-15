// Package day15 is our package for the 15th AoC day
package day15

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func getDistance(x1 int, y1 int, x2 int, y2 int) int {
	return int(math.Abs(float64(x1)-float64(x2)) + math.Abs(float64(y1)-float64(y2)))
}

func extractNumber(line string) int {
	line = strings.Replace(line, ",", "", -1)
	line = strings.Replace(line, ":", "", -1)
	numberAsString := strings.Split(line, "=")[1]
	number, _ := strconv.Atoi(numberAsString)
	return number
}

// RunPart1 is for the first star of the day
func RunPart1(input []string, example bool) int {
	pos := 2000000
	if example {
		pos = 10
	}
	occupied := make(map[int]struct{})
	noBeaconPossible := make(map[int]struct{})
	for _, line := range input {
		split := strings.Split(line, " ")
		sX := extractNumber(split[2])
		sY := extractNumber(split[3])
		bX := extractNumber(split[8])
		bY := extractNumber(split[9])

		if bY == pos {
			occupied[bX] = struct{}{}
		}

		distance := getDistance(sX, sY, bX, bY)

		for x := sX - distance; x < sX+distance; x++ {
			if _, exists := occupied[x]; exists {
				continue
			}
			if getDistance(sX, sY, x, pos) <= distance {
				noBeaconPossible[x] = struct{}{}
			}
		}
	}
	return len(noBeaconPossible)
}

// RunPart2 is for the second star of the day
func RunPart2(input []string, example bool) int {
	limit := 0
	if example {
		limit = 20
	} else {
		fmt.Println("day15 part 2 => wasn't able to solve part2 yet")
		return -1
	}
	var grid [20][20]bool
	for _, line := range input {
		split := strings.Split(line, " ")
		sX := extractNumber(split[2])
		sY := extractNumber(split[3])
		bX := extractNumber(split[8])
		bY := extractNumber(split[9])

		if sX >= 0 && sX < limit && sY >= 0 && sY < limit {
			grid[sX][sY] = true
		}
		if bX >= 0 && bX < limit && bY >= 0 && bY < limit {
			grid[bX][bY] = true
		}

		distance := getDistance(sX, sY, bX, bY)

		xMin := sX - distance
		if xMin < 0 {
			xMin = 0
		}

		xMax := sX + distance
		if xMax > limit {
			xMax = limit
		}

		yMin := sY - distance
		if yMin < 0 {
			yMin = 0
		}

		yMax := sY + distance
		if yMax > limit {
			yMax = limit
		}

		for x := sX - distance; x < sX+distance; x++ {
			for y := sY - distance; y < sY+distance; y++ {
				if x == sX && y == sY {
					continue
				}
				if x < xMin || x >= xMax || y < yMin || y >= yMax {
					continue
				}
				if grid[x][y] {
					continue
				}
				if getDistance(sX, sY, x, y) <= distance {
					grid[x][y] = true
				}
			}
		}
	}
	for x := 0; x < limit; x++ {
		for y := 0; y < limit; y++ {
			if !grid[x][y] {
				return x*4000000 + y
			}
		}
	}
	return -1
}
