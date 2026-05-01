func isAnagram(s string, t string) bool {
	letterMap := make(map[rune]int)

	for _, r := range s {
		if v, ok := letterMap[r]; ok {
			letterMap[r] = v + 1
		} else {
			letterMap[r] = 1
		}
	}

	for _, r := range t {
		if v, ok := letterMap[r]; ok {
			if v == 1 {
				delete(letterMap, r)
			} else {
				letterMap[r] = v - 1
			}
		} else {
			return false
		}
	}

	return len(letterMap) == 0
}
