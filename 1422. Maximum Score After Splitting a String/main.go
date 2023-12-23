package main

func maxScore(s string) int {
	length := len(s)
	count, max:= 0, 0

	for i := 1; i < length; i++ {
			for k := 0; k < length; k++ {
					if k < i {
							if s[k] == '0' {
									count++
							}
					} else if s[k] == '1' {
							count++
					}
			}

			if count > max { max = count }
			count = 0
	}

	return max
}