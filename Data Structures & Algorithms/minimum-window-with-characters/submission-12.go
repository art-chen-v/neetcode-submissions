func minWindow(s string, t string) string {
	// XYADBCXCA
	// ABCC
	l, r := 0, 0
	arr := [2]int{-1, -1}
	res := math.MaxInt

	counterS := make(map[byte]int)
	counterT := make(map[byte]int)
	sLen := 0

	for i := range t {
		counterT[t[i]]++
	}

	tLen := len(counterT)

	for r < len(s) {
		counterS[s[r]]++
		if _, ok := counterT[s[r]]; ok {
			if counterS[s[r]] == counterT[s[r]] {
				sLen++
			}
		}

		for sLen == tLen {
			if r - l < res {
				res = r - l
				arr[0], arr[1] = l, r
			}
			if _, ok := counterT[s[l]]; ok {
				counterS[s[l]]--
				if counterS[s[l]] < counterT[s[l]] {
					sLen--
				}
			}
			l++
		}
		r++
	}
	if arr == [2]int{-1, -1} {
		return ""
	}
	return s[arr[0]:arr[1] + 1]
}
