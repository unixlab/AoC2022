// Package day03 is our package for the 3rd AoC day
package day03

import (
	"strings"
)

func getCharNumber(char rune) int {
	charNumber := int(char) - 96
	if charNumber < 0 {
		return charNumber + 58
	}
	return charNumber
}

func isInList(lookup int, list []int) bool {
	for _, v := range list {
		if v == lookup {
			return true
		}
	}
	return false
}

func getCommonChars(str1 string, str2 string) []rune {
	var commonChars []rune
	for _, char := range str1 {
		for _, compChar := range str2 {
			if char == compChar {
				commonChars = append(commonChars, char)
			}
		}
	}
	return commonChars
}

// RunPart1 is for the first star of the day
func RunPart1(input []string) int {
	sum := 0
	for _, line := range input {
		indexHalf := len(line) / 2
		commonChar := getCommonChars(line[:indexHalf], line[indexHalf:])[0]
		charNumber := getCharNumber(commonChar)
		sum += charNumber
	}
	return sum
}

// RunPart2 is for the second star of the day
func RunPart2(input []string) int {
	sum := 0
	var lines []string
	for _, line := range input {
		lines = append(lines, line)
		if len(lines) == 3 {
			var commonChars []int
			for _, v := range getCommonChars(lines[0], lines[1]) {
				if !isInList(int(v), commonChars) {
					commonChars = append(commonChars, int(v))
				}
			}
			var commonString strings.Builder
			for _, v := range commonChars {
				commonString.WriteRune(rune(v))
			}
			commonChar := getCommonChars(commonString.String(), lines[2])[0]
			sum += getCharNumber(commonChar)
			lines = []string{}
		}
	}
	return sum
}
