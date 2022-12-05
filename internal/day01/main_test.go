package day01

import (
	"testing"

	"github.com/unixlab/AoC2022/internal/aoeinput"
)

func TestRunPart1Example(t *testing.T) {
	expectedResult := 24000
	actualResult := RunPart1(aoeinput.Read("../../", "day01", true))
	if actualResult != expectedResult {
		t.Fatalf("RunPart1() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart1(t *testing.T) {
	expectedResult := 67027
	actualResult := RunPart1(aoeinput.Read("../../", "day01", false))
	if actualResult != expectedResult {
		t.Fatalf("RunPart1() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2Example(t *testing.T) {
	expectedResult := 45000
	actualResult := RunPart2(aoeinput.Read("../../", "day01", true))
	if actualResult != expectedResult {
		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2(t *testing.T) {
	expectedResult := 197291
	actualResult := RunPart2(aoeinput.Read("../../", "day01", false))
	if actualResult != expectedResult {
		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
	}
}
