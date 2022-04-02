package main

import (
	"fmt"
)

func isPalindrome(str string) bool {
	for i := 0; i < len(str); i++ {
		j := len(str) - 1 - i
		if str[i] != str[j] {
			return false
		}
	}
	return true
}

func longestPalindrome2(s string) string {
	l := len(s)
	if l < 2  {
		return s
	}
	left := 0
	right := 0
	dp := make([][]int, 0)

	for i := 0; i < l; i++ {
		dp1 := make([]int, l)
		dp = append(dp, dp1)
		dp[i][i] = 1
	}

	// for i := 1; i < l; i++ {
	// 	for j := 0; j < i; j++ {
	// 		dp[i][j] = 0
	// 	}
	// }
	fmt.Println(dp)
	for j := 1; j < l; j++ {
		for i := 0; i < j; i++ {

			innerIsPalindrome := dp[i+1][j-1]==1 || (j-i) <= 2
			if s[i] ==  s[j] && innerIsPalindrome {
				dp[i][j] = 1
				if j-i > right -left {
					left = i
					right = j
				}
			}
		}
	}
	fmt.Println(dp)

	return s[left:right+1]

}

func longestPalindrome(s string) string {
	l := len(s)
	if l < 2  {
		return s
	}
	dp := make([][]string, 0)

	for i := 0; i < l; i++ {
		dp1 := make([]string, l)
		dp = append(dp, dp1)
		dp[i][i] = "T"
	}

	for i := 1; i < l; i++ {
		for j := 0; j < i; j++ {
			dp[i][j] = "F"
		}
	}
	// fmt.Println(dp)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if isPalindrome(s[i : j+1]) {
				dp[i][j] = "T"
			} else {
				dp[i][j] = "F"
			}

		}
	}
	// fmt.Println(dp)

	longestSubs := ""
	lenString := 0
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if dp[i][j] == "T" {
				str := s[i:j+1]
				if  len(str) > lenString {
					longestSubs = str
					lenString = len(str)
				}
			}
		}
	}

	return longestSubs

}

func main() {
	s := "babad"
	fmt.Println(longestPalindrome(s))
	fmt.Println(longestPalindrome2("bbcbbd"))

}
