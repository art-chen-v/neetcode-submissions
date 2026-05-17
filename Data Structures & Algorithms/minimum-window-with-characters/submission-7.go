func minWindow(s string, t string) string {
    window := make(map[byte]int)
	counterT := make(map[byte]int)
	res := [...]int{-1, -1}
	minLen := math.MaxInt
	for i := range t {
		counterT[t[i]]++
	}
	have, need := 0, len(counterT)
	l, r := 0, 0
	for r < len(s) {
		window[s[r]]++
		_, ok := counterT[s[r]]
		if ok && window[s[r]] == counterT[s[r]] {
			have++
		}
		for have == need {
			if (r-l+1) < minLen {
				minLen = r-l+1
				res[0], res[1] = l, r
			}
			window[s[l]]--
			_, ok := counterT[s[l]]
			if ok && window[s[l]] < counterT[s[l]] {
				have--
			}
			l++
		}
		r++
	}
	if res != [...]int{-1, -1} {
		return s[res[0]:res[1]+1]
	}
	return ""
}
