func characterReplacement(s string, k int) int {
	longest := 0
	count := make(map[byte]int)
	maxF := 0
	l := 0
	for r := 0; r < len(s); r++ {
		count[s[r]]++
		maxF = max(maxF, count[s[r]])
		for {
			if (r-l+1) - maxF <= k {
				break
			}
			count[s[l]]--
			l++
		}
		longest = r-l+1
	}
	return longest
}
