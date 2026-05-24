func characterReplacement(s string, k int) int {
	longest := 0
	counterF := make(map[byte]int)
	maxF := 0
	l := 0
	for r := 0; r < len(s); r++ {
		counterF[s[r]]++
		maxF = max(maxF, counterF[s[r]])
		for {
			if (r-l+1) - maxF <= k {
				longest = max(longest, r-l+1)
				break
			}
			counterF[s[l]]--
			l++
		}
	}
	return longest 
}
