package day05

import (
	"testing"

	"github.com/unixlab/AoC2022/internal/aoeinput"
)

func TestRunPart1Example(t *testing.T) {
	expectedResult := "CMZ"
	actualResult := RunPart1(aoeinput.Read("../../", "day05", true))
	if actualResult != expectedResult {
		t.Fatalf("RunPart1() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart1(t *testing.T) {
	expectedResult := "WHTLRMZRC"
	actualResult := RunPart1(aoeinput.Read("../../", "day05", false))
	if actualResult != expectedResult {
		t.Fatalf("RunPart1() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2Example(t *testing.T) {
	expectedResult := "MCD"
	actualResult := RunPart2(aoeinput.Read("../../", "day05", true))
	if actualResult != expectedResult {
		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2(t *testing.T) {
	expectedResult := "GMPMLWNMG"
	actualResult := RunPart2(aoeinput.Read("../../", "day05", false))
	if actualResult != expectedResult {
		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
	}
}
