// Package day13 is our package for the 13th AoC day
package day13

import (
	"encoding/json"
	"math"
	"reflect"
	"sort"
)

func min(a int, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func isInRightOrder(pkg1 interface{}, pkg2 interface{}) (bool, bool) {
	if reflect.TypeOf(pkg1).Kind() == reflect.Float64 && reflect.TypeOf(pkg2).Kind() == reflect.Float64 {
		if pkg1.(float64) == pkg2.(float64) {
			return true, false
		}
		if pkg1.(float64) < pkg2.(float64) {
			return true, true
		}
		return false, true
	} else if reflect.TypeOf(pkg1).Kind() == reflect.Slice && reflect.TypeOf(pkg2).Kind() == reflect.Slice {
		for i := 0; i < min(len(pkg1.([]interface{})), len(pkg2.([]interface{}))); i++ {
			res, abort := isInRightOrder(pkg1.([]interface{})[i], pkg2.([]interface{})[i])
			if abort {
				return res, true
			}
		}
		if len(pkg1.([]interface{})) > len(pkg2.([]interface{})) {
			return false, true
		}
		if len(pkg1.([]interface{})) < len(pkg2.([]interface{})) {
			return true, true
		}
	} else {
		if reflect.TypeOf(pkg1).Kind() != reflect.Slice {
			newPkg1 := make([]interface{}, 1)
			newPkg1[0] = pkg1
			pkg1 = newPkg1
		}
		if reflect.TypeOf(pkg2).Kind() != reflect.Slice {
			newPkg2 := make([]interface{}, 1)
			newPkg2[0] = pkg2
			pkg2 = newPkg2
		}
		return isInRightOrder(pkg1, pkg2)
	}
	return true, false
}

// RunPart1 is for the first star of the day
func RunPart1(input []string) int {
	var packets [][]interface{}
	for _, line := range input {
		if line == "" {
			continue
		}
		var packet []interface{}
		err := json.Unmarshal([]byte(line), &packet)
		if err != nil {
			return -1
		}
		packets = append(packets, packet)
	}
	sum := 0
	index := 1
	for i := 0; i < len(packets); i += 2 {
		order, _ := isInRightOrder(packets[i], packets[i+1])
		if order {
			sum += index
		}
		index++
	}
	return sum
}

// RunPart2 is for the second star of the day
func RunPart2(input []string) int {
	var packets [][]interface{}
	input = append(input, "[[2]]")
	input = append(input, "[[6]]")
	for _, line := range input {
		if line == "" {
			continue
		}
		var packet []interface{}
		err := json.Unmarshal([]byte(line), &packet)
		if err != nil {
			return -1
		}
		packets = append(packets, packet)
	}
	sort.Slice(packets, func(i, j int) bool {
		res, _ := isInRightOrder(packets[i], packets[j])
		return res
	})
	sum := 1
	for idx, packet := range packets {
		jsonBytes, _ := json.Marshal(packet)
		packetString := string(jsonBytes)
		if packetString == "[[2]]" || packetString == "[[6]]" {
			sum *= idx + 1
		}
	}
	return sum
}
