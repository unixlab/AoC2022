// Package day12 is our package for the 12th AoC day
package day12

import (
	"strings"

	"github.com/RyanCarrier/dijkstra"
)

// Run takes the part number and can be used for both parts
func Run(input []string, part int) int {
	var grid [][]int
	var startIdx, destIdx int
	var startingPoints []int
	idx := 0
	graph := dijkstra.NewGraph()
	for _, line := range input {
		var row []int
		for _, char := range strings.Split(line, "") {
			elevation := char[0] - 97
			if elevation == 242 {
				elevation = 0
				startIdx = idx
			}
			if elevation == 228 {
				elevation = 25
				destIdx = idx
			}
			if elevation == 0 {
				startingPoints = append(startingPoints, idx)
			}
			graph.AddVertex(idx)
			idx++
			row = append(row, int(elevation))
		}
		grid = append(grid, row)
	}
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			for i := y - 1; i <= y+1; i++ {
				if i < 0 || i >= len(grid) {
					continue
				}
				for j := x - 1; j <= x+1; j++ {
					if j < 0 || j >= len(grid[i]) {
						continue
					}
					if i == y && j == x {
						continue
					}
					if y != i && x != j {
						continue
					}
					if grid[y][x]+1 >= grid[i][j] {
						node1 := y*len(grid[0]) + x
						node2 := i*len(grid[0]) + j
						graph.AddArc(node1, node2, 1)
					}
				}
			}
		}
	}
	if part == 1 {
		path, _ := graph.Shortest(startIdx, destIdx)
		return int(path.Distance)
	}
	var minDistance int64
	minDistance = -1
	for _, point := range startingPoints {
		path, err := graph.Shortest(point, destIdx)
		if err != nil {
			continue
		}
		if path.Distance < minDistance || minDistance == -1 {
			minDistance = path.Distance
		}
	}
	return int(minDistance)
}
