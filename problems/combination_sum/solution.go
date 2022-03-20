package main

import (
	"fmt"
)

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	if len(candidates) == 0 {
		return res
	}
	
	res = solutionFinder(candidates, target, 0, 0, res, []int{})
	
	return res
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

func solutionFinder(candidates []int, target, sum, index int, res [][]int, curComb []int) [][]int {
	if index > len(candidates) {
		return res
	}
	
	if sum > target {
		return res
	} else if sum == target {
		res = append(res, copy(curComb))
		return res
	}

	for i := index; i < len(candidates); i++ {
		curComb = append(curComb, candidates[i])
		res = solutionFinder(candidates, target, sum + candidates[i], i, res, curComb)
		curComb = remove(curComb, candidates[i])
	}
	

	return res
}

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 18
	fmt.Println("result", combinationSum(candidates, target))
}
