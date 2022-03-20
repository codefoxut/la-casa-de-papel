package main

import (
	"fmt"
)

func climbStairs(n int) int {
    dp := []int{
		0, 1, 2,
	}

	for i:=3; i<=n; i++ {
		dp = append(dp, dp[i-1]+dp[i-2])
	}
	return dp[n]
	
}


func main() {
	fmt.Println("climb 3", climbStairs(3))
	fmt.Println("climb 4", climbStairs(4))
	fmt.Println("climb 5", climbStairs(5))
}