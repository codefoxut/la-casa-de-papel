package main

import (
	"fmt"
	"sort"
)

func containDuplicateBrute(nums []int) bool {
	var flag bool
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}

	return flag
}

func containDuplicateSorting(nums []int) bool {
	var flag bool
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}

	return flag
}

func containDuplicateMaps(nums []int) bool {
	var flag bool
	mapOfArray := make(map[int]bool, len(nums))
	for _, item := range nums {
		if val := mapOfArray[item]; val {
			return true
		}
		mapOfArray[item] = true
	}

	return flag
}

func containDuplicateSets(nums []int) bool {
	mapOfArray := make(map[int]bool)
	for _, item := range nums {
		mapOfArray[item] = true
	}
	return len(nums) != len(mapOfArray)
}

func main() {
	arr := []int{1, 4, 6, 7, 20, 5, 2, 11, 15, 3, 5}
	fmt.Printf("Founnd duplicates 1 %v\n", containDuplicateBrute(arr))
	fmt.Printf("Founnd duplicates 2 %v\n", containDuplicateSorting(arr))
	fmt.Printf("Founnd duplicates 3 %v\n", containDuplicateMaps(arr))
	fmt.Printf("Founnd duplicates 4 %v\n", containDuplicateSets(arr))

}
