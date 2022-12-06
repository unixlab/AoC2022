// Package day06 is our package for the 6th AoC day
package day06

func allCharsUniq(input string) bool {
	var store uint32
	for _, char := range input {
		charNumber := char - 97
		if store&(1<<charNumber) > 0 {
			return false
		}
		store |= 1 << charNumber
	}
	return true
}

// Run is for both stars of the day
func Run(input string, window int) int {
	for index := 0; index < len(input)-window; index++ {
		if allCharsUniq(input[index : index+window]) {
			return index + window
		}
	}
	return -1
}
