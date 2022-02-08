package main

import (
	"fmt"
)

func valid_mountain_array(arr []int) (ret bool) {

	decrementTracker :=  -1
	incrementTracker :=  -1


	for i, item := range arr {
		if i == 0 {
			incrementTracker = i
			continue
		}

		if item > arr[i-1]  && decrementTracker > 0 {
			ret = false
			break
		} else if item > arr[i-1] {
			incrementTracker = i
		} else {
			decrementTracker = i
		}
	}
	if decrementTracker == len(arr) - 1  && incrementTracker > 0 {
		ret = true
	}
	return

}

func valid_mountain_array2(arr []int) (ret bool) {
	//  INCREASING subsequence
	var tracker int

	for i:=1; i<len(arr) && arr[i] > arr[i-1]; i++ {
		tracker = i
	}

	if tracker ==  0 || tracker  ==  len(arr) -1 {
		return false
	}

	for i := tracker+1 ; i <  len(arr)  && arr[i] < arr[i-1]; i++ {
		
		tracker = i
	}
	if tracker ==  len(arr) -1 {
		return true
	}
	return 
}


func main()  {
	arr1 := []int{1, 3, 5, 8, 10, 13, 12, 8, 4, 2}
	fmt.Println(arr1,  valid_mountain_array(arr1))
	arr2 := []int{1, 3, 5, 8, 10, 13, 12, 15, 4, 2}
	fmt.Println(arr2,  valid_mountain_array(arr2))
	arr3 := []int{1, 3, 5, 8, 10, 13}
	fmt.Println(arr3,  valid_mountain_array(arr3))
	arr4 := []int{15, 13, 12, 8, 4, 2}
	fmt.Println(arr4,  valid_mountain_array(arr4))
	arr5 := []int{1, 3, 5, 8, 10, 13, 11}
	fmt.Println(arr5,  valid_mountain_array(arr5))


}