package main

func searchInsert(nums []int, target int) int {
	indexPoint := 0
	for index, value := range nums {
			if value == target {
					return index
			}

			if !(index + 1 > len(nums)) {
					if value < target {
							indexPoint = index + 1
					}
			}
	}
	return indexPoint
}