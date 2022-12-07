// Package day07 is our package for the 7th AoC day
package day07

import (
	"sort"
	"strconv"
	"strings"
)

func calculateDirectories(input []string) map[string]int {
	directories := make(map[string]int, 0)
	currentPath := ""
	for _, line := range input {
		if strings.HasPrefix(line, "$ ls") {
			continue
		}
		if strings.HasPrefix(line, "dir") {
			continue
		}
		if strings.HasPrefix(line, "$ cd") {
			relPath := strings.Split(line, " ")[2]
			if relPath == ".." {
				idx := strings.LastIndex(currentPath, "/")
				currentPath = currentPath[:idx]
			} else if relPath == "/" {
				currentPath = "/"
			} else {
				if !strings.HasSuffix(currentPath, "/") {
					currentPath += "/"
				}
				currentPath += relPath
			}
			continue
		}
		size, _ := strconv.Atoi(strings.Split(line, " ")[0])
		tree := strings.Split(currentPath, "/")
		for i := len(tree); i > 0; i-- {
			if currentPath == "/" {
				break
			}
			path := strings.Join(tree[:i], "/")
			if path == "" {
				continue
			}
			directories[path] += size
		}
		directories["/"] += size
	}
	return directories
}

// RunPart1 is for the first star of the day
func RunPart1(input []string) int {
	sum := 0
	for _, dir := range calculateDirectories(input) {
		if dir < 100000 {
			sum += dir
		}
	}
	return sum
}

// RunPart2 is for the second star of the day
func RunPart2(input []string) int {
	directories := calculateDirectories(input)
	spaceNeeded := 3e7 - (7e7 - directories["/"])
	var sizes []int
	for _, dirSize := range directories {
		sizes = append(sizes, dirSize)
	}
	sort.Ints(sizes)
	for _, size := range sizes {
		if size >= spaceNeeded {
			return size
		}
	}
	return 0
}
