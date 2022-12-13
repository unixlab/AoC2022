package day13

import (
	"testing"

	"github.com/unixlab/AoC2022/internal/aoeinput"
)

func TestRunPart1Example(t *testing.T) {
	expectedResult := 13
	actualResult := RunPart1(aoeinput.Read("../../", "day13", true))
	if actualResult != expectedResult {
		t.Fatalf("RunPart1() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart1(t *testing.T) {
	expectedResult := 6187
	actualResult := RunPart1(aoeinput.Read("../../", "day13", false))
	if actualResult != expectedResult {
		t.Fatalf("RunPart1() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart1JsonError(t *testing.T) {
	expectedResult := -1
	actualResult := RunPart1([]string{"ThisIsNotJson"})
	if actualResult != expectedResult {
		t.Fatalf("RunPart1() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2Example(t *testing.T) {
	expectedResult := 140
	actualResult := RunPart2(aoeinput.Read("../../", "day13", true))
	if actualResult != expectedResult {
		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2(t *testing.T) {
	expectedResult := 23520
	actualResult := RunPart2(aoeinput.Read("../../", "day13", false))
	if actualResult != expectedResult {
		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2JsonError(t *testing.T) {
	expectedResult := -1
	actualResult := RunPart2([]string{"ThisIsNotJson"})
	if actualResult != expectedResult {
		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
	}
}
