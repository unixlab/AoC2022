// Package aoeinput handles all the AoC input
package aoeinput

import (
	"bufio"
	"os"
	"strings"
)

// Read is our global function to read AoC input
func Read(day string, example bool) []string {
	var path strings.Builder
	if example {
		path.WriteString("examples")
	} else {
		path.WriteString("inputs")
	}
	path.WriteString("/")
	path.WriteString(day)
	file, err := os.Open(path.String())
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
