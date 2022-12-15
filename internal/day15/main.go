// Package day15 is our package for the 15th AoC day
package day15

import (
	"math"
	"strconv"
	"strings"
)

func absDiff(x int, y int) int {
	return int(math.Abs(float64(x) - float64(y)))
}

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

// Entry is one entry from our input
type Entry struct {
	Sensor    Pos
	Beacon    Pos
	Distance  int
	YDistance int
}

// Pos represents one position in the grid
type Pos struct {
	X int
	Y int
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
	limit := 4000000
	if example {
		limit = 20
	}
	var entries []Entry
	occupied := make(map[Pos]bool)
	for _, line := range input {
		var newEntry Entry
		split := strings.Split(line, " ")
		newEntry.Sensor.X = extractNumber(split[2])
		newEntry.Sensor.Y = extractNumber(split[3])
		occupied[newEntry.Sensor] = true
		newEntry.Beacon.X = extractNumber(split[8])
		newEntry.Beacon.Y = extractNumber(split[9])
		occupied[newEntry.Beacon] = true
		newEntry.Distance = getDistance(
			newEntry.Sensor.X,
			newEntry.Sensor.Y,
			newEntry.Beacon.X,
			newEntry.Beacon.Y,
		)
		newEntry.YDistance = absDiff(newEntry.Sensor.Y, newEntry.Beacon.Y)
		entries = append(entries, newEntry)
	}
	for x := 0; x < limit; x++ {
	nextPos:
		for y := 0; y < limit; y++ {
			pos := Pos{x, y}
			if !occupied[pos] {
				for _, e := range entries {
					distance := getDistance(
						e.Sensor.X,
						e.Sensor.Y,
						x,
						y,
					)
					if distance <= e.Distance {
						distance = e.Distance - absDiff(e.Sensor.X, x)
						y += distance + (e.Sensor.Y - y)
						continue nextPos
					}
				}
				return x*4e6 + y
			}
		}
	}
	return -1
}
