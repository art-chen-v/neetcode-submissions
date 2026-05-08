func characterReplacement(s string, k int) int {
	l, r := 0, 0
	longest := 0
	counter := make(map[byte]int)
// s="A ABAB BA"
// k=1
// 4
	for r < len(s) {
		counter[s[r]]++
		for {
			if (r-l+1) - mostFreqLetterCount(counter) <= k {
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

func mostFreqLetterCount(counter map[byte]int) int {
	m := 0
	for _, v := range counter {
		m = max(m, v)
	}
	return m
}
