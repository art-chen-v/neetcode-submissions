func lengthOfLongestSubstring(s string) int {
	// abcabcbb
	longest := 0
	set := make(map[byte]struct{})
	l, r := 0, 0
	for r < len(s) {
		for {
			_, ok := set[s[r]]
			if !ok {
				break
			}
			delete(set, s[l])
			l++
		}
		set[s[r]] = struct{}{}
		longest = max(longest, r-l+1)
		r++
	}
	return longest
}
