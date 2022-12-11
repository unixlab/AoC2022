package day11

import (
	"testing"

	"github.com/unixlab/AoC2022/internal/aoeinput"
)

func TestRunPart1Example(t *testing.T) {
	expectedResult := 10605
	actualResult := Run(aoeinput.Read("../../", "day11", true), 1)
	if actualResult != expectedResult {
		t.Fatalf("RunPart1() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart1(t *testing.T) {
	expectedResult := 102399
	actualResult := Run(aoeinput.Read("../../", "day11", false), 1)
	if actualResult != expectedResult {
		t.Fatalf("RunPart1() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2Example(t *testing.T) {
	expectedResult := 2713310158
	actualResult := Run(aoeinput.Read("../../", "day11", true), 2)
	if actualResult != expectedResult {
		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2(t *testing.T) {
	expectedResult := 23641658401
	actualResult := Run(aoeinput.Read("../../", "day11", false), 2)
	if actualResult != expectedResult {
		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunUnknownInput(t *testing.T) {
	expectedResult := -1
	actualResult := Run([]string{"ThisIsWrongInputData"}, 1)
	if actualResult != expectedResult {
		t.Fatalf("RunPart1() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunUnknownOp(t *testing.T) {
	expectedResult := -1
	actualResult := Run([]string{"Operation: new = old / 6"}, 1)
	if actualResult != expectedResult {
		t.Fatalf("RunPart1() = expected %v, got %v", expectedResult, actualResult)
	}
}
