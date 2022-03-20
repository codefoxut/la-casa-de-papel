package main

import (
	"fmt"
	"math"
	"sort"
)

func minArray(a []int) int {
	minValue := int(math.MaxInt32)
	for _, x := range a {
		if minValue > x {
			minValue = x
		}
	}
	return minValue
}

func coinChange1(coins []int, amount int) int {
	sort.Ints(coins)
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	minCoins := make([]int, len(coins))
	for i, coin := range coins {
		val := coinChange1(coins, amount-coin)
		if val < 0 {
			val = int(math.MaxInt32)
		} else {
			val = 1 + val
		}
		minCoins[i] = val
	}
	v := minArray(minCoins)
	if int(math.MaxInt32) == v {
		return -1
	}
	return v
}

func coinChange2(coins []int, amount int) int {
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	sort.Sort(sort.Reverse(sort.IntSlice(coins)))

	coinCount, flag := solutionFinder(coins, amount, 0)
	if flag {
		return coinCount
	}
	return -1
}

func solutionFinder(coins []int, amount int, coinCount int) (int, bool) {
	if amount < 0 {
		return -1, false
	}
	if amount == 0 {
		return coinCount, true
	}
	for i := 0; i < len(coins); i++ {
		count, val := solutionFinder(coins, amount-coins[i], coinCount+1)
		fmt.Println(val)
		if val {
			return count, true
		}
	}
	return -1, false
}

func contains(s []int, num int) bool {
	for _, v := range s {
		if v == num {
			return true
		}
	}

	return false
}

func coinChangeLoop(coins []int, amount int) int {
	if amount <= 0 {
		return 0
	}

	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}

	for i := 1; i <= amount; i++ {
		var values []int
		values = append(values, dp[i])
		for _, coin := range coins {
			if i >= coin {
				values = append(values, dp[i-coin]+1)
			}
		}
		dp[i] = minArray(values)

	}
	fmt.Println(dp)

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func coinChange(coins []int, amount int) int {
	if amount <= 0 {
		return 0
	}

	dp := make([]int, amount+1)

	for _, coin := range coins {
		if coin <= amount {
			dp[coin] = 1
		}
	}

	for i := 1; i <= amount; i++ {
		var values []int
		var val int
		if !contains(coins, i) {
			for _, c := range coins {
				if i> c {
					val = dp[i-c]
					values = append(values, val+1)
				}
			}
			if len(values) == 0 {
				dp[i] = int(math.MaxInt32)
			} else {
				dp[i] = minArray(values)
			}
		}
	}
	// fmt.Println(dp)

	if dp[amount] == int(math.MaxInt32) {
		return -1
	}
	return dp[amount]
}

func main() {
	coins := []int{1, 2, 5}
	amount := 11
	fmt.Println(coinChange(coins, amount))
	fmt.Println(coinChangeLoop(coins, amount))

}
