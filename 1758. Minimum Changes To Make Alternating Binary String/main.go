package main

func minOperations(s string) int {
	zero, one := 0, 0
	for i, ch := range s {
		if i%2 == 0 && ch == '0' || i%2 == 1 && ch == '1' {
			one++
		} else {
			zero++
		}
	}
	return min(zero, one)
}