package main

import (
	"fmt"
)


func moveZero1(arr []int){
	var nonZeroCounter int

	for _, elem := range arr {
		if elem != 0 {
			arr[nonZeroCounter] = elem
			nonZeroCounter += 1
		}
	}

	for ;nonZeroCounter < len(arr); nonZeroCounter++   {
		arr[nonZeroCounter] = 0
		 
	}

	fmt.Println(arr)
}

func moveZero(arr []int){
	var nonZeroCounter int

	for i, elem := range arr {
		if elem != 0 {
			arr[i] = 0
			arr[nonZeroCounter] = elem
			nonZeroCounter += 1
		}
	}

	fmt.Println(arr)
}

func main(){
	data := []int{0,  1, 0, 3, 0,  15, 0,  4, 100, 0,  10}

	moveZero(data)
}