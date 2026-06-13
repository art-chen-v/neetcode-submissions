func minWindow(s string, t string) string {
	res := [2]int{-1, -1}
	shortest := math.MaxInt
	window := make(map[byte]int)
	counterT := make(map[byte]int)
	for i := range t {
		counterT[t[i]]++
	}
	have, need := 0, len(counterT)
	l, r := 0, 0

	for r < len(s) {
		
		if _, ok := counterT[s[r]]; ok {
			window[s[r]]++
			if window[s[r]] == counterT[s[r]] {
				have++
			}
		}


		for have == need {
			if (r - l) < shortest {
				shortest = r - l
				res = [...]int{l, r}
			}
			
			if _, ok := counterT[s[l]]; ok {
				window[s[l]]--
				if window[s[l]] < counterT[s[l]] {
					have--
				}
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
