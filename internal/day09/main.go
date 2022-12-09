// Package day09 is our package for the 9th AoC day
package day09

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func getDistance(x int, y int) int {
	return int(math.Abs(math.Abs(float64(x)) - math.Abs(float64(y))))
}

// invertedMaxOne converts any positive number to -1, any negative number to 1 and keeps 0
func invertedMaxOne(input int) int {
	if input > 0 {
		return -1
	} else if input < 0 {
		return 1
	}
	return 0
}

// Run takes the number of knots can be used for both parts
func Run(input []string, numberOfKnots int) int {
	xPos := make([]int, numberOfKnots)
	yPos := make([]int, numberOfKnots)
	for knot := 0; knot < numberOfKnots; knot++ {
		xPos[knot] = 500
		yPos[knot] = 500
	}
	visited := make(map[string]int, 0)
	for _, line := range input {
		direction := strings.Split(line, " ")[0]
		steps, _ := strconv.Atoi(strings.Split(line, " ")[1])
		for i := 0; i < steps; i++ {
			switch direction {
			case "L":
				xPos[0]--
			case "R":
				xPos[0]++
			case "U":
				yPos[0]++
			case "D":
				yPos[0]--
			}
			for knot := 1; knot < numberOfKnots; knot++ {
				distanceY := getDistance(yPos[knot], yPos[knot-1])
				distanceX := getDistance(xPos[knot], xPos[knot-1])
				if distanceY > 1 || distanceX > 1 {
					yPos[knot] += invertedMaxOne(yPos[knot] - yPos[knot-1])
					xPos[knot] += invertedMaxOne(xPos[knot] - xPos[knot-1])
				}
			}
			coordinate := fmt.Sprintf("%dx%d", yPos[numberOfKnots-1], xPos[numberOfKnots-1])
			visited[coordinate]++
		}
	}
	return len(visited)
}
