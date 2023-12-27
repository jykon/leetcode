package main

func minCost(colors string, neededTime []int) int {
	time := 0
	for i:= 0; i + 1 < len(neededTime); i++ {
			if colors[i] != colors[i+1] {
					continue
			}

			if (neededTime[i] <= neededTime[i+1]) {
					time += neededTime[i]
			} else {
					time += neededTime[i+1]
					neededTime[i+1] = neededTime[i]
			}
	}

	return time
}
