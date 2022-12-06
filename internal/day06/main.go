// Package day06 is our package for the 6th AoC day
package day06

// allCharsUnique checks if all characters (a-z) in the input are unique
func allCharsUnique(input string) bool {
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

// Run takes the window size and those can be used for both parts
func Run(input string, window int) int {
	for index := 0; index < len(input)-window; index++ {
		if allCharsUnique(input[index : index+window]) {
			return index + window
		}
	}
	return -1
}
