package day09

import (
	"testing"

	"github.com/unixlab/AoC2022/internal/aoeinput"
)

func TestRunPart1Example(t *testing.T) {
	expectedResult := 13
	actualResult := Run(aoeinput.Read("../../", "day09", true), 2)
	if actualResult != expectedResult {
		t.Fatalf("RunPart1() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart1(t *testing.T) {
	expectedResult := 5883
	actualResult := Run(aoeinput.Read("../../", "day09", false), 2)
	if actualResult != expectedResult {
		t.Fatalf("RunPart1() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2Example1(t *testing.T) {
	expectedResult := 1
	actualResult := Run(aoeinput.Read("../../", "day09", true), 10)
	if actualResult != expectedResult {
		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2Example2(t *testing.T) {
	expectedResult := 36
	actualResult := Run(aoeinput.Read("../../", "day09p2", true), 10)
	if actualResult != expectedResult {
		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
	}
}

func TestRunPart2(t *testing.T) {
	expectedResult := 2367
	actualResult := Run(aoeinput.Read("../../", "day09", false), 10)
	if actualResult != expectedResult {
		t.Fatalf("RunPart2() = expected %v, got %v", expectedResult, actualResult)
	}
}
