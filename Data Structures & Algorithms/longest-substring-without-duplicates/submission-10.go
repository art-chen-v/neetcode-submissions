func lengthOfLongestSubstring(s string) int {
	// abcabcbb
	longest := 0
	set := make(map[byte]struct{})
	l := 0
	for r := 0; r < len(s); r++ {
		for {
			if _, ok := set[s[r]]; !ok {
				break
			}
			delete(set, s[l])
			l++
		}
		longest = max(longest, r-l+1)
		set[s[r]] = struct{}{}
	}
	return longest
}
