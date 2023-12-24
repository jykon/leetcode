package main

import "math"

func maxProductDifference(nums []int) int {
	low1, low2 := math.MaxInt, math.MaxInt
	great1, great2 := math.MinInt, math.MinInt

	for _, v := range nums {
			if great1 < v {
					great2, great1 = great1, v
			} else if great2 < v {
					great2 = v
			}
			if low1 > v {
					low2, low1 = low1, v
			} else if low2 > v {
					low2 = v
			}
	}

	return great1*great2-low1*low2
}