func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sMap := make(map[rune]int)

	for _, r := range s {
		sMap[r]++
	}

	for _, r := range t {
		count, ok := sMap[r]
		if !ok || count == 0 {
			return false
		}
		sMap[r]--
	}

	return true
}
