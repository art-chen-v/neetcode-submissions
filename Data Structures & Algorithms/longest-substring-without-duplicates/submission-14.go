func lengthOfLongestSubstring(s string) int {
	longest := 0
	set := make(map[byte]struct{})
	l := 0
	for r := 0; r < len(s); r++ {
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
	}
	return longest
}
