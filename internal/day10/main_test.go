package day10

import (
	"strings"
	"testing"

	"github.com/unixlab/AoC2022/internal/aoeinput"
)

func TestRunPart1Example(t *testing.T) {
	expectedResult := 13140
	actualResult, _ := Run(aoeinput.Read("../../", "day10", true), false)
	if actualResult != expectedResult {
		t.Fatalf("Run() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart1(t *testing.T) {
	expectedResult := 14320
	actualResult, _ := Run(aoeinput.Read("../../", "day10", false), false)
	if actualResult != expectedResult {
		t.Fatalf("Run() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2Example(t *testing.T) {
	expectedResult := `##..##..##..##..##..##..##..##..##..##..
###...###...###...###...###...###...###.
####....####....####....####....####....
#####.....#####.....#####.....#####.....
######......######......######......####
#######.......#######.......#######.....
`
	_, actualResult := Run(aoeinput.Read("../../", "day10", true), false)
	if actualResult != expectedResult {
		t.Fatalf("Run() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2(t *testing.T) {
	expectedResult := `###...##..###..###..#..#..##..###....##.
#..#.#..#.#..#.#..#.#.#..#..#.#..#....#.
#..#.#....#..#.###..##...#..#.#..#....#.
###..#....###..#..#.#.#..####.###.....#.
#....#..#.#....#..#.#.#..#..#.#....#..#.
#.....##..#....###..#..#.#..#.#.....##..
`
	_, actualResult := Run(aoeinput.Read("../../", "day10", false), false)
	if actualResult != expectedResult {
		t.Fatalf("Run() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunInputValidation(t *testing.T) {
	expectedResult1 := -1
	expectedResult2 := ""
	actualResult1, actualResult2 := Run([]string{"addy 3"}, false)
	if actualResult1 != expectedResult1 {
		t.Fatalf("Run() = expected %v, got %v", expectedResult1, actualResult1)
	}
	if actualResult2 != expectedResult2 {
		t.Fatalf("Run() = expected %v, got %v", expectedResult2, actualResult2)
	}
}

func TestRunColor(t *testing.T) {
	input := make([]string, 250)
	for i := 0; i < len(input); i++ {
		input[i] = "noop"
	}
	expectedResult1 := 720
	var expectedResult2Builder strings.Builder
	for i := 0; i < 6; i++ {
		for y := 0; y < 40; y++ {
			expectedResult2Builder.WriteString(" ")
		}
		expectedResult2Builder.WriteString("\n")
	}
	actualResult1, actualResult2 := Run(input, true)
	if actualResult1 != expectedResult1 {
		t.Fatalf("Run() = expected %v, got %v", expectedResult1, actualResult1)
	}
	if actualResult2 != expectedResult2Builder.String() {
		t.Fatalf("Run() = expected %v, got %v", expectedResult2Builder.String(), actualResult2)
	}
}
