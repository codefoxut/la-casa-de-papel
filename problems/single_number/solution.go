package main

import (
	"errors"
	"fmt"
)

func brute_force(arr []int) (int, error) {
	numMap := make(map[int]int, len(arr))

	for _, item := range arr {
		val := numMap[item]
		numMap[item] = val + 1
	}

	for item, val := range numMap {
		if val == 1 {
			return item, nil
		}
	}
	return 0, errors.New("item not found")
}

func using_set(arr []int) int {
	sumOfTheArray := 0
	mapOfArray := make(map[int]bool, len(arr))

	for _, item := range arr {
		sumOfTheArray += item
		mapOfArray[item] = true
	}

	sumOfMap := 0
	for item := range mapOfArray {
		sumOfMap += item
	}
	
	return (2 * sumOfMap) - sumOfTheArray
}

func  using_xor(arr []int) int {
	sumOfArray := 0 

	for _, item := range arr {
		sumOfArray =  sumOfArray ^ item
	}
	return sumOfArray
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 5, 4, 2, 3, 1, 6}
	v, err := brute_force(arr)
	if err == nil {
		fmt.Printf("missing element in the array is %d\n", v)
	} else {
		fmt.Printf("error with solution  is %v\n", err)
	}

	fmt.Printf("missing element in the array using set is %d\n", using_set(arr))

	fmt.Printf("missing element in the array usign xor is %d\n", using_xor(arr))

}
