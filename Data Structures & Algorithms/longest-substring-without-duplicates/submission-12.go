func lengthOfLongestSubstring(s string) int {
	longest := 0
	counter := make(map[byte]struct{})
	l := 0
	for r := 0; r < len(s); r++ {
		for {
			if _, ok := counter[s[r]]; !ok {
				break
			}
			delete(counter, s[l])
			l++
		}
		counter[s[r]] = struct{}{}
		longest = max(longest, r - l + 1)
	}
	return longest
}
