package main

import (
	"fmt"
	"strings"
)

func  max(a, b int)  int {
	if a > b {
		return a
	}  
	return b
}

func longest_substring(str string) (int, string) {
	var p1, p2, ans  int
	var ch rune
	var longestStr string
	mapOfCharaters := make(map[rune]int)
	runeSample := []rune(str)

	for p1 < len(str) && p2 < len(str) {
		ch = runeSample[p2]
		if value, ok := mapOfCharaters[ch]; ok {
			p1 = max(p1, value+1)
		} 
		mapOfCharaters[ch] = p2
		if p2  - p1 + 1  > ans {
			ans = max(ans, p2-p1+1)
			longestStr = str[p1:p2+1]
		}
		

		p2++
		
	}
	fmt.Println(mapOfCharaters)
	return ans, longestStr
}

func main(){
	s1 := "abcabcbb"
	fmt.Println(longest_substring(s1))

	s2 :=  "pwwkew"
	fmt.Println(longest_substring(s2))

	s3 := strings.ToLower("ArunerepresentsaUnicodePoint")
	fmt.Println(longest_substring(s3))

	s4 :=  "abaacad"
	fmt.Println(longest_substring(s4))

}