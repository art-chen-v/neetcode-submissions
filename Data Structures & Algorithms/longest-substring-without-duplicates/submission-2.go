func lengthOfLongestSubstring(s string) int {

	// "dvadf"
	// pwwkew
	m := make(map[uint8]struct{})
	longest := 0
	counter := 0
	l := 0
	i := 0
	for i < len(s) {
		if _, ok := m[s[i]]; ok {
			longest = max(counter, longest)
			m = make(map[uint8]struct{})
			counter = 0
			l++
			i = l
		} else {
			m[s[i]] = struct{}{}
			counter++
			i++
		}
	}

	return max(longest, counter)
}
