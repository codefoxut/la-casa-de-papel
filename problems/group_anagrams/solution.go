package main

import (
	"fmt"
	"sort"
)

func findHash(s string) string {
	runeSlice := []rune(s)
	sort.Slice(runeSlice, func(i, j int) bool {
		return runeSlice[i] < runeSlice[j]
	})
	return string(runeSlice)
}

func groupAnagrams(strs []string) [][]string {
	hashMaps := make(map[string][]string)
	soln := make([][]string, 0)
	for _, item := range strs {
		h := findHash(item)
		if val, ok := hashMaps[h]; ok {
			hashMaps[h] = append(val, item)
		} else {
			hashMaps[h] = []string{item}
		}
	}
	for _, v := range hashMaps {
		soln = append(soln, v)
	}
	return soln

}

func main() {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Printf("output %v\n", groupAnagrams(input))

}
