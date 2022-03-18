package main

import (
	"fmt"
)

func minWindow(s string, t string) string {
	// create map of pattern
	mapPattern := make(map[rune]int)
	for _, ch := range t {
		mapPattern[ch] = mapPattern[ch] + 1
	}

	fmt.Printf("%v\n", mapPattern)

	mapString := make(map[rune]int)
	runeString := []rune(s)
	var left, right int
	minSpan := []int{len(s) + 1, 0, 0}
	if len(s) > 0 {
		mapString[runeString[left]] = 1
	}

	for left < len(s) && right < len(s) {
		// if string has the pattern, we will increment left
		//  if string did not have thee pattern, we will increment right.
		fmt.Printf("string %v\n", mapString)
		if checkMapsEqualty(mapPattern, mapString) {
			length := right - left + 1
			if length < minSpan[0] {
				minSpan[0] = length
				minSpan[1] = left
				minSpan[2] = right
			}
			fmt.Printf("minspan %v\n", minSpan)
			// remove left element from mapString
			mapString[runeString[left]] -= 1
			if mapString[runeString[left]] <= 0 {
				delete(mapString, runeString[left])
			}
			// increment left
			left++

		} else {
			right++
			if right < len(s) {
				mapString[runeString[right]] = mapString[runeString[right]] + 1
			}

		}

	}

	if minSpan[0] <= len(s) {
		return string(runeString[minSpan[1] : minSpan[2]+1])
	}

	return ""
}

func checkMapsEqualty(mapPattern, mapString map[rune]int) bool {
	for k, v1 := range mapPattern {
		if v2, ok := mapString[k]; !ok || v2 < v1 { // minimum
			return false
		}
	}
	if len(mapPattern) == 0 {
		return false
	}
	return true
}

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"

	// fmt.Println("output 1 is", minWindow(s, t))

	s = "a"
	t = "aa"
	fmt.Println("output 2 is", minWindow(s, t))
}
