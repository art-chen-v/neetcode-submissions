func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	letterMap := make(map[rune]int)

	for _, r := range s {
		letterMap[r]++
	}

	for _, r := range t {
		count, ok := letterMap[r]
		if !ok || count == 0 {
			return false
		}
		letterMap[r]--
	}

	return true
}
