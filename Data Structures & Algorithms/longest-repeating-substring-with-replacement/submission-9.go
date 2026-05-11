func characterReplacement(s string, k int) int {
	longest := 0
	counter := make(map[byte]int)
	l := 0
	maxF := 0
	for r := 0; r < len(s); r++ {
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
	}
	return longest
}
