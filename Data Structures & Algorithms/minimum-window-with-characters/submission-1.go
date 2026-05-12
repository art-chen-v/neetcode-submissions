type result struct {
	l int
	r int
}

func minWindow(s string, t string) string {
    if t == "" {
		return ""
	}
	countT := make(map[byte]int)
	window := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		countT[t[i]]++
	}
	have := 0
	need := len(countT)
	res := result{}
	l := 0
	minLen := math.MaxInt
	for r := 0; r < len(s); r++ {
		window[s[r]]++
		_, ok := countT[s[r]]
		if ok && window[s[r]] == countT[s[r]] {
			have++
		}

		for have == need {
			if (r-l+1) < minLen {
				minLen = r-l+1
				res = result{
					l: l,
					r: r,
				}
			}
			window[s[l]]--
			_, ok := countT[s[l]]
			if ok && window[s[l]] < countT[s[l]] {
				have--
			}
			l++
		}
	}
	if minLen < math.MaxInt {
		return s[res.l:res.r+1]
	}
	return ""
}
