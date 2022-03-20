package main

import (
	"fmt"
	"strings"
)

func getMappedLetters(char rune) []string {
	letterMap := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}
	return letterMap[string(char)]
}

func copy(s []string) []string {
	t := make([]string, 0, len(s))
	t = append(t, s...)
	return t
}
func remove(s []string, num string) []string {
	index := -1
	for i, v := range s {
		if v == num {
			index = i
		}
	}

	return append(s[:index], s[index+1:]...)
}

func letterCombinations(digits string) []string {
	listOfLetters := make([][]string, 0, len(digits))
	for _, item := range digits {
		listOfLetters = append(listOfLetters, getMappedLetters(item))
	}
	fmt.Printf("list of letters %v\n", listOfLetters)

	result := make([]string, 0)
	result = solutionFinder(listOfLetters, result, []string{}, 0)
	

	return result
}

func solutionFinder(listOfLetters [][]string, result []string, cur []string, index int) []string {
	fmt.Println(cur, index, result)
	if index > len(listOfLetters) {
		return result
	}
	if len(cur) == len(listOfLetters) {
		if len(cur) > 0 {
			result = append(result, strings.Join(cur,  ""))
		}
		fmt.Println(result, cur)
		return result
	}

	for i := 0; i < len(listOfLetters[index]); i++ {
		ch := listOfLetters[index][i]
		cur = append(cur, ch)
		result = solutionFinder(listOfLetters, result, cur, index+1)
		cur = remove(cur, ch)
	}
	return result
}

func main() {
	digits := ""
	fmt.Println(letterCombinations(digits))
}
