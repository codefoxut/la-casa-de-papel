package main

import (
	"fmt"
)



func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(height))
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func trap(height []int) int {
	left := make([]int, len(height))
	right := make([]int, len(height))

	for index, item := range height {
		if index  == 0 {
			left[index] = item
			continue
		}
		left[index] = max(left[index-1], item)
	}

	for j := len(height) - 1; j >=0; j-- {
		if j == len(height) - 1 {
			right[j] = height[j]
			continue
		}
		right[j] = max(right[j+1], height[j])
	} 

	// fmt.Printf("%v\t%v\n", left, right)
	total  := 0
	for i:= range height {
		total  +=  min(right[i], left[i]) - height[i]
		

	}
	return total
}
