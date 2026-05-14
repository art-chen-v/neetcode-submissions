func minWindow(s string, t string) string {
    minLen := math.MaxInt
	res := [...]int{-1, -1}
	window := make(map[byte]int)
	counterT := make(map[byte]int)
	for i := range t {
		counterT[t[i]]++
	}
	have, need := 0, len(counterT)
	l := 0

	for r := range s {
		window[s[r]]++
		if _, ok := counterT[s[r]]; ok && window[s[r]] == counterT[s[r]] {
			have++
		}
		for have == need {
			if (r - l + 1) < minLen {
				minLen = r - l + 1
				res[0], res[1] = l, r
			}
			window[s[l]]--
			if _, ok := counterT[s[l]]; ok && window[s[l]] < counterT[s[l]] {
				have--
			}
			l++
		}
	}
	if res != [...]int{-1, -1} {
		return s[res[0]:res[1] + 1]
	}
	return ""
}
