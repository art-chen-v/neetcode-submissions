func lengthOfLongestSubstring(s string) int {
	set := make(map[byte]struct{})
	longest := 0
	l := 0
	for r := 0; r < len(s); r++ {
		for {
			if _, ok := set[s[r]]; !ok {
				break
			}
			delete(set, s[l])
			l++
		}
		set[s[r]] = struct{}{}
		longest = max(longest, r-l+1)
	}
	return longest
}
