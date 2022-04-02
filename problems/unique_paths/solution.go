package main

import (
	"fmt"
)

func uniquePaths(m int, n int) int {
	dp := make([][]int, 0)

	for i := 0; i < m; i++ {
		dp1 := make([]int, n)
		dp = append(dp, dp1)
	}
	// fmt.Println(dp)

	//  first row can be reached only one  way
	for i:=0; i<m; i++ {
		dp[i][0] = 1
	}
	// first column can be reached only one way.
	for i:=0;i<n; i++  {
		dp[0][i] = 1
	}

	for i:=1;i<m;  i++ {
		for j:=1;j<n;j++  {
			dp[i][j] = dp[i-1][j]+dp[i][j-1]
		}
	}
	// fmt.Println(dp)
	return dp[m-1][n-1]
}

func main() {
	m := 3
	n := 7
	fmt.Println(uniquePaths(m, n))
}
