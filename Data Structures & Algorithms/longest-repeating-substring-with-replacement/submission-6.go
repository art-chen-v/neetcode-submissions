func characterReplacement(s string, k int) int {
	longest := 0
	counter := make(map[byte]int)

	l := 0
	for r := 0; r < len(s); r++ {
		counter[s[r]]++
		for {
			if (r-l+1) - maxLetterCounter(counter) <= k {
				break
			}
			counter[s[l]]--
			l++
		}
		longest = max(longest, r-l+1)
	}
	return longest
}

func maxLetterCounter(counter map[byte]int) int {
	maxCounter := 0
	for _, v := range counter {
		maxCounter = max(maxCounter, v)
	}
	return maxCounter
}