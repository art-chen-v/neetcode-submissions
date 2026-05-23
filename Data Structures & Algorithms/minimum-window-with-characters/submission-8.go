func minWindow(s string, t string) string {
    counterT := make(map[byte]int)
	for i := range t {
		counterT[t[i]]++
	}
	window := make(map[byte]int)
	have, need := 0, len(counterT)
	l, r := 0, 0
	minCount := math.MaxInt
	res := [2]int{-1, -1}

	for ; r < len(s); r++ {
		window[s[r]]++
		_, ok := counterT[s[r]]
		if ok && window[s[r]] == counterT[s[r]] {
			have++
		}
		for have == need {
			if (r-l+1) < minCount {
				minCount = r-l+1
				res[0], res[1] = l, r
			}
			window[s[l]]--
			_, ok := counterT[s[l]]
			if ok && window[s[l]] < counterT[s[l]] {
				have--
			}
			l++
		}
	}
	if res == [...]int{-1, -1} {
		return ""
	}
	return s[res[0]:res[1]+1]
}
