package main

import (
	"fmt"
	"sort"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func missing_number(num []int) int  {
	// find the sum of numbers in the array.
	sum := 0
	for _, item := range num {
		sum += item
	}

	expectedSum := (len(num)  * (len(num) + 1)) / 2

	return expectedSum - sum

}

func missing_numbers_with_sorting(nums []int) []int {
	sort.Ints(nums)

	missingNums := make([]int, 0)
	countMissingNums := 0

	for index, value := range nums  {
		for index + countMissingNums != value  {
			missingNums = append(missingNums, index+countMissingNums)
			countMissingNums++
		}
	}

	return missingNums
}

func missing_numbers(nums []int) []int {
	numMaps := make(map[int]bool)
	maxOfNums := 0
	for _, item := range nums {
		numMaps[item] = true
		maxOfNums = max(maxOfNums, item)
	}

	missingNums := make([]int, 0)

	for i :=0; i<=maxOfNums; i++ {
		if ok := numMaps[i]; !ok {
			missingNums = append(missingNums, i)
		}
	}

	return missingNums
}

func main(){
	fmt.Println(4*5/2, 5*4/2)

	arr := []int{1, 2, 3, 5, 6, 7, 8,  9, 10, 0}

	fmt.Println(missing_number(arr))

	fmt.Println(missing_numbers(arr))

	arr2 :=  []int{1, 10}
	fmt.Println(missing_numbers(arr2))

}