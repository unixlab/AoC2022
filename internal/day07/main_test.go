package day07

import (
	"testing"

	"github.com/unixlab/AoC2022/internal/aoeinput"
)

func TestRunPart1Example(t *testing.T) {
	expectedResult := 95437
	actualResult := RunPart1(aoeinput.Read("../../", "day07", true))
	if actualResult != expectedResult {
		t.Fatalf("RunPart1() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart1(t *testing.T) {
	expectedResult := 1915606
	actualResult := RunPart1(aoeinput.Read("../../", "day07", false))
	if actualResult != expectedResult {
		t.Fatalf("RunPart1() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2Example(t *testing.T) {
	expectedResult := 24933642
	actualResult := RunPart2(aoeinput.Read("../../", "day07", true))
	if actualResult != expectedResult {
		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2(t *testing.T) {
	expectedResult := 5025657
	actualResult := RunPart2(aoeinput.Read("../../", "day07", false))
	if actualResult != expectedResult {
		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
	}
}
