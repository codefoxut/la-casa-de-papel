package main

import (
	"fmt"
)


func isBadVersion(n int) bool {
	return n > 25
}


func findBadVersion(arr []int) int {
	left := 0
	right := len(arr) - 1

	for left < right {
		mid := (left + right) / 2
		// fmt.Println(left, right, mid)
		if isBadVersion(arr[mid]) {
			right = mid
		} else {
			left = mid +  1
		}
	}
	return arr[left]
}

func main(){
	arr :=  []int{1, 18, 25, 29,30, 45, 49, 50, 79, 100}

	fmt.Println(findBadVersion(arr))

}