func characterReplacement(s string, k int) int {
	// AAABABB
	counter := make(map[byte]int)
	longest := 0
	l, r := 0, 0
	maxF := 0
	for r < len(s) {
		counter[s[r]]++
		maxF = max(maxF, counter[s[r]])
		for {
			if (r-l+1) - maxF <= k {
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
