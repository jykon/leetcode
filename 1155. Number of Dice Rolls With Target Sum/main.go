package main


func numRollsToTarget(n int, k int, target int) int {
	cur := make([]int, target + 1)
	for i := 1; i <= k && i <= target; i++ {
			cur[i] = 1
	}

	for x := 2; x <= n; x++ {
			next := make([]int, target + 1)
			for i := x; i <= x * k && i <= target; i++ {
					for j := i - 1; j >= i - k && j >= 0; j-- {
							next[i] = (next[i] + cur[j]) % (1e9 + 7)
					}
			}
			cur = next
	}
	return cur[target]
}