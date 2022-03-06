package main

import  (

	"fmt"
)

func majorityElement(nums []int) int {
	mapOfNums := make(map[int]int, len(nums))

	for _, item := range nums {
		_, ok := mapOfNums[item]
		if ok {
			mapOfNums[item] += 1
		} else {
			mapOfNums[item] = 1
		}
	}
	for k, v := range mapOfNums {
		if v > len(nums)/2 {
			return k
		}
	}
	return 0

}

func main() {
	arr := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Printf("Majority element is %v \n",  majorityElement(arr))

}
