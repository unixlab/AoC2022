package day15

import (
	"testing"

	"github.com/unixlab/AoC2022/internal/aoeinput"
)

func TestRunPart1Example(t *testing.T) {
	expectedResult := 26
	actualResult := RunPart1(aoeinput.Read("../../", "day15", true), true)
	if actualResult != expectedResult {
		t.Fatalf("RunPart1() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart1(t *testing.T) {
	expectedResult := 5181556
	actualResult := RunPart1(aoeinput.Read("../../", "day15", false), false)
	if actualResult != expectedResult {
		t.Fatalf("RunPart1() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2Example(t *testing.T) {
	expectedResult := 56000011
	actualResult := RunPart2(aoeinput.Read("../../", "day15", true), true)
	if actualResult != expectedResult {
		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2ExampleNoFree(t *testing.T) {
	expectedResult := -1
	actualResult := RunPart2([]string{"0 0 x=0 y=0 0 0 0 0 x=19 y=19"}, true)
	if actualResult != expectedResult {
		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2(t *testing.T) {
	expectedResult := 12817603219131
	actualResult := RunPart2(aoeinput.Read("../../", "day15", false), false)
	if actualResult != expectedResult {
		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
	}
}
