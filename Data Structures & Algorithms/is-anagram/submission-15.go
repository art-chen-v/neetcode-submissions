func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sM := make(map[byte]int)

	for i := range s {
		sM[s[i]]++
	}

// racecar
// carracr

	for i := range t {
		if _, ok := sM[t[i]]; ok {
			sM[t[i]]--
			if sM[t[i]] == 0 {
				delete(sM, t[i])
			}
		} else {
			return false
		}
	}
	if len(sM) == 0 {
		return true
	}
	return false
}
