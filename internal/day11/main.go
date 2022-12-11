// Package day11 is our package for the 11th AoC day
package day11

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

// Monkey is the struct for the part1 input
type Monkey struct {
	Items          []int
	Operation      func(old int, sec int) int
	SecondNumber   int
	Test           int
	NewMonkeyTrue  int
	NewMonkeyFalse int
	Inspections    int
}

func getBored(input int) int {
	return int(math.Floor(float64(input) / 3))
}

func greatestCommonDivisor(num1 int, num2 int) int {
	for num2 > 0 {
		store := num2
		num2 = num1 % num2
		num1 = store
	}
	return num1
}

func leastCommonMultiplier(num1 int, num2 int) int {
	return num1 * num2 / greatestCommonDivisor(num1, num2)
}

// Run takes the part number and can be used for both parts
func Run(input []string, part int) int {
	var monkeys []Monkey
	var newMonkey Monkey
	part2Divisor := 1
	for _, line := range input {
		if strings.HasPrefix(line, "Monkey") {
			continue
		}
		line = strings.TrimLeft(line, " ")
		if strings.HasPrefix(line, "Starting") {
			line = strings.Replace(line, " ", "", -1)
			itemsString := strings.Split(line, ":")[1]
			for _, item := range strings.Split(itemsString, ",") {
				newItem, _ := strconv.Atoi(item)
				newMonkey.Items = append(newMonkey.Items, newItem)
			}
			continue
		}
		if strings.HasPrefix(line, "Operation") {
			op := strings.Replace(strings.Split(line, "=")[1], " ", "", -1)
			if op == "old*old" {
				number := 0
				newMonkey.SecondNumber = number
				newMonkey.Operation = func(old int, sec int) int { return old * old }
			} else if strings.HasPrefix(op, "old+") {
				numberAsString := strings.Split(op, "+")[1]
				number, _ := strconv.Atoi(numberAsString)
				newMonkey.SecondNumber = number
				newMonkey.Operation = func(old int, sec int) int { return old + sec }
			} else if strings.HasPrefix(op, "old*") {
				numberAsString := strings.Split(op, "*")[1]
				number, _ := strconv.Atoi(numberAsString)
				newMonkey.SecondNumber = number
				newMonkey.Operation = func(old int, sec int) int { return old * sec }
			} else {
				return -1
			}
			continue
		}
		if strings.HasPrefix(line, "Test") {
			testValString := strings.Split(line, "by")[1]
			testValString = strings.Trim(testValString, " ")
			newMonkey.Test, _ = strconv.Atoi(testValString)
			part2Divisor = leastCommonMultiplier(part2Divisor, newMonkey.Test)
			continue
		}
		if strings.HasPrefix(line, "If true") {
			testValString := strings.Split(line, "to monkey")[1]
			testValString = strings.Trim(testValString, " ")
			newMonkey.NewMonkeyTrue, _ = strconv.Atoi(testValString)
			continue
		}
		if strings.HasPrefix(line, "If false") {
			testValString := strings.Split(line, "to monkey")[1]
			testValString = strings.Trim(testValString, " ")
			newMonkey.NewMonkeyFalse, _ = strconv.Atoi(testValString)
			continue
		}
		if line == "" {
			monkeys = append(monkeys, newMonkey)
			newMonkey = Monkey{}
			continue
		}
		return -1
	}
	monkeys = append(monkeys, newMonkey)
	maxRounds := 20
	if part == 2 {
		maxRounds = 1e4
	}
	for round := 0; round < maxRounds; round++ {
		for k, monkey := range monkeys {
			for _, item := range monkey.Items {
				monkeys[k].Inspections++
				item = monkey.Operation(item, monkey.SecondNumber)
				if part == 1 {
					item = getBored(item)
				} else {
					item = item % part2Divisor
				}
				if item%monkey.Test == 0 {
					monkeys[monkey.NewMonkeyTrue].Items = append(monkeys[monkey.NewMonkeyTrue].Items, item)
					monkeys[k].Items = monkeys[k].Items[1:]
				} else {
					monkeys[monkey.NewMonkeyFalse].Items = append(monkeys[monkey.NewMonkeyFalse].Items, item)
					monkeys[k].Items = monkeys[k].Items[1:]
				}
			}
		}
	}
	var inspections []int
	for _, monkey := range monkeys {
		inspections = append(inspections, monkey.Inspections)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(inspections)))
	return inspections[0] * inspections[1]
}
