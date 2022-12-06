package day06

import (
	"testing"

	"github.com/unixlab/AoC2022/internal/aoeinput"
)

func TestRunWindow4NoMatch(t *testing.T) {
	expectedResults := []int{-1, -1, -1}
	for k, v := range []string{"aaa", "aaaa", "aaaaa"} {
		actualResult := Run(v, 4)
		if actualResult != expectedResults[k] {
			t.Fatalf("Run(window = 4) = expected %v, got %v", expectedResults[k], actualResult)
		}
	}
}

func TestRunWindow4Example(t *testing.T) {
	expectedResults := []int{7, 5, 6, 10, 11}
	for k, v := range aoeinput.Read("../../", "day06", true) {
		actualResult := Run(v, 4)
		if actualResult != expectedResults[k] {
			t.Fatalf("Run(window = 4) = expected %v, got %v", expectedResults[k], actualResult)
		}
	}
}

func TestRunWindow4(t *testing.T) {
	expectedResults := []int{1480}
	for k, v := range aoeinput.Read("../../", "day06", false) {
		actualResult := Run(v, 4)
		if actualResult != expectedResults[k] {
			t.Fatalf("Run(window = 4) = expected %v, got %v", expectedResults[k], actualResult)
		}
	}
}

func TestRunWindow14Example(t *testing.T) {
	expectedResults := []int{19, 23, 23, 29, 26}
	for k, v := range aoeinput.Read("../../", "day06", true) {
		actualResult := Run(v, 14)
		if actualResult != expectedResults[k] {
			t.Fatalf("Run(window = 14) = expected %v, got %v", expectedResults[k], actualResult)
		}
	}
}

func TestRunWindow14(t *testing.T) {
	expectedResults := []int{2746}
	for k, v := range aoeinput.Read("../../", "day06", false) {
		actualResult := Run(v, 14)
		if actualResult != expectedResults[k] {
			t.Fatalf("Run(window = 14) = expected %v, got %v", expectedResults[k], actualResult)
		}
	}
}
