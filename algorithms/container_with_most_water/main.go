package main

import  (
	"fmt"
)

type solution struct {
	area int
	p1 int
	p2 int
}



func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func brute_force(arr []int) solution {

	sol :=  solution{
		area:0,
		p1:0,
		p2:0,
	}
	

	for i:=0; i<len(arr) -1; i++  {
		for j:= i+1; j < len(arr); j++ {
			area := min(arr[i], arr[j]) * (j-i)
			if area > sol.area {
				sol.area = area
				sol.p1 = i
				sol.p2 = j
			}
		}
	}
	return sol
}

func intuitive(arr []int) solution {
	l := 0
	r := len(arr) - 1

	sol :=  solution{
		area:0,
		p1:0,
		p2:0,
	}

	for l < r {
		area := min(arr[l], arr[r]) *  (r-l)

		// fmt.Println(arr[l], arr[r], area)

		if area >  sol.area {
			sol.area = area
			sol.p1 = l
			sol.p2 = r
		}

		if arr[l] >  arr[r] {
			// fmt.Println("1", arr[r], arr[r-1], r,  l)
			for l < r {
				r -= 1
				if arr[r+1] < arr[r] {
					break
				} 				
				// fmt.Println(arr[r], arr[r-1], r)
			}
		} else {
			// fmt.Println("2", arr[l], arr[l+1], l,  r)
			for (l < r) {
				l += 1
				if arr[l-1] < arr[l] {
				 break
			}  
				// fmt.Println(arr[l], arr[l+1], l)
			}
		}
	} 
	return sol
}


func main() {
	arr := []int{5, 2, 8, 4, 3, 6, 4, 2, 7, 9}
	fmt.Println(arr, brute_force(arr))


	fmt.Println(arr, intuitive(arr))

	arr2 := []int{5, 9, 2, 4}
	fmt.Println(arr2, brute_force(arr2))


	fmt.Println(arr2, intuitive(arr2))
}