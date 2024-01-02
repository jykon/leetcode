func maxLengthBetweenEqualCharacters(s string) int {
	max := 0
	if len(s) == 2 && s[0] == s[1] {
			return 0
	}

	for i, v := range s {
			counter := strings.Count(s, string(v))

			if counter > 1 {
					lastIndex := strings.LastIndex(s, string(v))
					if lastIndex != -1 {
							if lastIndex - i > max {
									max = lastIndex - i
							}
					}
			}

	}
	if max == 0 {
			return -1
	}
	return max - 1
}