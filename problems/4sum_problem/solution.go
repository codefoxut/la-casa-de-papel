package main

import (
	"fmt"
)

func fourSumCountBrute(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	sum := 0
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			for k := 0; k < len(nums3); k++ {
				for l := 0; l < len(nums4); l++ {
					if nums1[i]+nums2[j]+nums3[k]+nums4[l] == 0 {
						sum++
					}
				}
			}
		}
	}
	return sum
}

func fourSumCountMaps(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	map12 := make(map[int]int)
	map34 := make(map[int]int)
	count := 0
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			sum := nums1[i] + nums2[j]
			if _, ok1 := map12[sum]; ok1 {
				map12[sum]++
			} else {
				map12[sum] = 1
			}
		}
	}
	for k := 0; k < len(nums3); k++ {
		for l := 0; l < len(nums4); l++ {
			sum := nums3[k] + nums4[l]
			if _, ok2 := map34[sum]; ok2 {
				map34[sum]++
			} else {
				map34[sum] = 1
			}
		}
	}

	for k, v := range map12 {
		target := 0 - k
		if num, ok := map34[target]; ok {
			count += (v * num)
		}
	}
	return count
}

func fourSumCountMaps2(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	map12 := make(map[int]int)
	count := 0
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			sum := nums1[i] + nums2[j]
			if _, ok1 := map12[sum]; ok1 {
				map12[sum]++
			} else {
				map12[sum] = 1
			}
		}
	}
	for k := 0; k < len(nums3); k++ {
		for l := 0; l < len(nums4); l++ {
			target := 0 - (nums3[k] + nums4[l])
			if num, ok := map12[target]; ok {
				count += num
			}
		}
	}

	return count
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{-2, -1}
	nums3 := []int{-1, 2}
	nums4 := []int{0, 2}
	fmt.Printf("solution %d\n", fourSumCountBrute(nums1, nums2, nums3, nums4))

	fmt.Printf("solution 2: %d\n", fourSumCountMaps(nums1, nums2, nums3, nums4))

	fmt.Printf("solution 3: %d\n", fourSumCountMaps2(nums1, nums2, nums3, nums4))
}
