package main

import (
	"fmt"
	"sort"
)

func contains(s []int, num int) bool {
	for _, v := range s {
		if v == num {
			return true
		}
	}

	return false
}
func copy(s []int) []int {
	t := make([]int, 0, len(s))
	t = append(t, s...)
	return t
}

func remove(s []int, num int) []int {
	index := -1
	for i, v := range s {
		if v == num {
			index = i
		}
	}

	return append(s[:index], s[index+1:]...)
}

func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	sort.Ints(nums)
	return solutionFinder(nums, result, nil, 0)

}

func solutionFinder(nums []int, result [][]int, cur []int, index int) [][]int {
	if cur == nil {
		cur = []int{}
	}
	if index >= len(nums) {
		return result
	}
	fmt.Printf("%v\n", cur)
	result = append(result, copy(cur))
	for i := index; i < len(nums); i++ {
		if !contains(cur, nums[i]) {
			cur = append(cur, nums[i])
			result = solutionFinder(nums, result, cur, i)
			cur = remove(cur, nums[i])
		}
	}
	return result
}

func main() {
	arr := []int{1, 2, 3}
	fmt.Printf("%v\n", subsets(arr))

}
