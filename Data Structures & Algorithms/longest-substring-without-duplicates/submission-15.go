func lengthOfLongestSubstring(s string) int {
	counter := make(map[byte]struct{})
	res := 0
	l, r := 0, 0
	for r < len(s) {
		for {
			if _, ok := counter[s[r]]; !ok {
				res = max(res, r-l+1)
				break
			}
			delete(counter, s[l])
			l++
		}
		counter[s[r]] = struct{}{}
		r++
	}
	return res
}
