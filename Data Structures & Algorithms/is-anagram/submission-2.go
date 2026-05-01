func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	letterMap := make(map[rune]int)

	for _, r := range s {
		letterMap[r]++
	}

	for _, r := range t {
		if count, ok := letterMap[r]; ok {
			if count == 0 {
				return false
			} else {
				letterMap[r]--
			}
		} else {
			return false
		}
	}

	return true
}
