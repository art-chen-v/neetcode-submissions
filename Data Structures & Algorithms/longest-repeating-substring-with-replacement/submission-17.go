func characterReplacement(s string, k int) int {
	// s = "AAABABB", k = 1
	longest := 0
	counter := make(map[byte]int)
	l, r := 0, 0
	freq := 0
	for r < len(s) {
		counter[s[r]]++
		freq = max(freq, counter[s[r]])
		for {
			if (r-l+1) - freq <= k {
				longest = max(longest, r-l+1)
				break
			}
			counter[s[l]]--
			l++
		}
		r++
	}
	return longest
}
