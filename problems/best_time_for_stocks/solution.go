package main

import (
	"fmt"
	"math"
)

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {
	buyingPrice := int(math.MaxInt32)
	profit :=  0

	for _, price := range prices {
		if buyingPrice > price {
			buyingPrice = price
		}  else {
			profit  = max(profit, price - buyingPrice)
		}
	}

	return profit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}
