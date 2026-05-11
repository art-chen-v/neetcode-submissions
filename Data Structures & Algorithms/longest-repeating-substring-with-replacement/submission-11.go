func characterReplacement(s string, k int) int {
// AAABAB B
	longest := 0
	counter := make(map[byte]int)
	l := 0
	for r := 0; r < len(s); r++ {
		counter[s[r]]++
		for {
			if (r-l+1) - maxLetterCount(counter) <= k {
				break
			}
			counter[s[l]]--
			l++
		}
		longest = max(longest, r-l+1)
	}
	return longest
}

func maxLetterCount(counter map[byte]int) int {
	longest := 0
	for _, v := range counter {
		longest = max(longest, v)
	}
	return longest
}
