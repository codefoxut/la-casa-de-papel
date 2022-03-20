package main

import (
	"fmt"
)

func isPalindrome(str string) bool {
	for i := 0; i < len(str); i++ {
		j := len(str) - 1 - i
		if str[i] != str[j] {
			return false
		}
	}
	return true
}

// func allPartitionsPalindrome(strSet [][]rune) bool {
// 	for i := 0; i < len(strSet); i++ {
// 		if !isPalindrome(strSet[i]) {
// 			return false
// 		}
// 	}
// 	return true
// }

func partition(s string) [][]string {
	// sRune := []rune(s)
	return solutionFinder(s, [][]string{}, []string{})
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

func solutionFinder(s string, result [][]string, cur []string) [][]string {
	// fmt.Println("start", cur, s == "", s)
	if len(cur) > 0 && s == "" {
		result = append(result, copy(cur))
		// fmt.Println(result, "result")
		return result
	}

	for i := 1; i <= len(s); i++ {
		segment := s[:i]
		// fmt.Println(segment, "Segments", s[i:])
		if isPalindrome(segment) {
			cur = append(cur, segment)
			result = solutionFinder(s[i:], result, cur)
			cur = remove(cur, segment)
		}
	}

	return result
}

func main() {
	s := "aaaa"
	fmt.Println(partition(s))
}
