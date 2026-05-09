func characterReplacement(s string, k int) int {
	longest := 0
	counter := make(map[byte]int)
	maxf := 0
	l := 0
	for r := 0; r < len(s); r++ {
		counter[s[r]]++
		maxf = max(maxf, counter[s[r]])
		for (r-l+1) - maxf > k {
			counter[s[l]]--
			l++
		}
		longest = max(longest, r-l+1)
	}
	return longest
}
