package main

import (
	"errors"
	"fmt"
)

func twoSumBrute(arr []int, target int) (int, int, error) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				return i, j, nil
			}
		}
	}
	return 0, 0, errors.New("element not found")
}

func twoSumMap(arr []int, target int) (int, int, error) {
	mapOfArray := make(map[int]int, len(arr))
	for index, item := range arr {
		mapOfArray[item] = index
		if target-item == item {
			continue
		}
		if val, ok := mapOfArray[target-item]; ok {
			return val, index, nil
		}

	}

	return 0, 0, errors.New("element not found")

}

func main() {
	arr := []int{12, 11, 4, 15, 9, 13}
	sum := 24
	x, y, err := twoSumBrute(arr, sum)
	if err == nil {
		fmt.Printf("order for required sum %d is %d, %d\n", sum, x, y)
	} else {
		fmt.Printf("error %v \n", err)
	}

	sum = 22
	x, y, err = twoSumMap(arr, sum)
	if err == nil {
		fmt.Printf("order for required sum %d is %d, %d\n", sum, x, y)
	} else {
		fmt.Printf("error %v \n", err)
	}

}
