package main

import (
	"fmt"
	"strconv"
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func addBinary(binaryA, binaryB string) string {
	var soln string
	i := len(binaryA)
	j := len(binaryB)
	a := []rune(binaryA)
	b := []rune(binaryB)
	sol := make([]string, max(i, j)+1)
	lengthOfSol := len(sol)
	var sum, carry int
	for i > 0 || j > 0 || carry > 0 {
		var x, y int
		if i > 0 {
			x, _ = strconv.Atoi(string(a[i-1]))
			i--
		}
		if j > 0 {
			y, _ = strconv.Atoi(string(b[j-1]))
			j--
		}
		sum = carry + x + y

		sol[lengthOfSol-1] = strconv.Itoa(sum % 2)
		carry = sum / 2

		lengthOfSol--
	}

	if sol[0] == "0" {
		soln = strings.Join(sol[1:], "")
	} else {
		soln = strings.Join(sol, "")
	}

	return soln
}

func main() {
	fmt.Println(addBinary("1111", "1"))
}
