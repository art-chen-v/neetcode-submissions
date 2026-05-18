func characterReplacement(s string, k int) int {
	counter := make(map[byte]int)
	longest := 0
	l, r := 0, 0
	for r < len(s) {
		counter[s[r]]++
		for {
			if (r-l+1) - getMaxLetterCount(counter) <= k {
				break
			}
			counter[s[l]]--
			l++
		}
		longest = max(longest, r-l+1)
		r++
	}
	return longest
}

func getMaxLetterCount(counter map[byte]int) int {
	m := 0
	for _, v := range counter {
		m = max(m, v)
	}
	return m
}
