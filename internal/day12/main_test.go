package day12

import (
	"testing"

	"github.com/unixlab/AoC2022/internal/aoeinput"
)

func TestRunPart1Example(t *testing.T) {
	expectedResult := 31
	actualResult := Run(aoeinput.Read("../../", "day12", true), 1)
	if actualResult != expectedResult {
		t.Fatalf("Run() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart1(t *testing.T) {
	expectedResult := 380
	actualResult := Run(aoeinput.Read("../../", "day12", false), 1)
	if actualResult != expectedResult {
		t.Fatalf("Run() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2Example(t *testing.T) {
	expectedResult := 29
	actualResult := Run(aoeinput.Read("../../", "day12", true), 2)
	if actualResult != expectedResult {
		t.Fatalf("Run() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2(t *testing.T) {
	expectedResult := 375
	actualResult := Run(aoeinput.Read("../../", "day12", false), 2)
	if actualResult != expectedResult {
		t.Fatalf("Run() = expected %v, got %v", expectedResult, actualResult)
	}
}
