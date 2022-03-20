package main

import (
	"fmt"
)

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func rob(nums []int) int {

	sol := [][]int{
		{nums[0], 0},
	}
	fmt.Println(sol)
	var  rob, notRob int
	for i := 1; i < len(nums); i++ {
		//if rob
		rob = sol[i-1][1] + nums[i]
		notRob = max(sol[i-1][0], sol[i-1][1])
		fmt.Println("house",  i, nums[i], rob, notRob)
		sol = append(sol, []int{rob,  notRob})
	}
	return max(sol[len(nums)-1][0], sol[len(nums)-1][1])
}

func main() {
	nums := []int{2,7,9,3,1}
	fmt.Println(rob(nums))
}
