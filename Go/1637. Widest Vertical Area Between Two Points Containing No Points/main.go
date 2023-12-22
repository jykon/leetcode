package main

import (
	"cmp"
	"slices"
)

func maxWidthOfVerticalArea(points [][]int) int {
	slices.SortFunc(points, func(a, b []int) int {
		return cmp.Compare(a[0], b[0])
	})

	max_gap := 0
	for prev_idx, curr := range points[1:] {
		max_gap = max(max_gap, curr[0] - points[prev_idx][0])
	}

	return max_gap
}