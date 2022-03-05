package main

import (
	"fmt"
)


func findPositions(arr []int,  target int) (int, int) {
	if arr[0] > target || arr[len(arr)-1] < target {
		return -1, -1
	}
	leftIndex, rightIndex := findSearchSpan(arr, target)

	if leftIndex == rightIndex {
		if arr[leftIndex] != target {
			return -1 , -1
		} else {
			return leftIndex, rightIndex
		}
	}

	fmt.Println("Span  found", leftIndex, rightIndex)


	return  findLeft(arr, target, leftIndex, rightIndex), findRight(arr, target, leftIndex, rightIndex)
}

func findSearchSpan(arr []int, target int) (int, int) {
	var left, right, mid int
	right = len(arr) - 1
	
	for left < right {
		fmt.Println(left, right, "span")
		mid  = (left + right) / 2
		if arr[mid] == target {
			return left, right
		} else if arr[mid] > target {
			right = mid-1
		}  else {
			left = mid + 1
		}
	}
	return left, right
}

func findLeft(arr []int, target, leftIndex, rightIndex int) int {
	var mid int
	for leftIndex < rightIndex {
		fmt.Println(leftIndex, rightIndex, "left")
		mid  = (leftIndex + rightIndex) / 2
		if arr[mid] > target {
			rightIndex =  mid - 1 
		} else if arr[mid] < target {
			leftIndex = mid + 1
		}  else {
			rightIndex = mid
		}
	}
	return  leftIndex
}

func findRight(arr []int, target, leftIndex, rightIndex int) int {
	var mid int
	for leftIndex < rightIndex {
		fmt.Println(leftIndex, rightIndex, "right")
		mid  = (leftIndex + rightIndex) / 2
		if arr[mid] < target {
			leftIndex =  mid + 1
		} else if arr[mid] > target {
			rightIndex = mid - 1
		} else {
			leftIndex  = mid
		}
	}
	return  leftIndex
}

func main() {

	arr1 := []int{10, 11, 11, 11, 11, 14, 15}
	fmt.Println(findPositions(arr1, 11))

	fmt.Println(findPositions(arr1, 15))

	fmt.Println(findPositions(arr1, 13))

}