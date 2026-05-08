func characterReplacement(s string, k int) int {
	longest := 0
	counter := make(map[byte]int)
	l, r := 0, 0
// AABABBA
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
	maxCount := 0
	for _, v := range counter {
		maxCount = max(maxCount, v)
	}
	return maxCount
}
