func isAnagram(s string, t string) bool {
	sMap := make(map[rune]int)
	tMap := make(map[rune]int)



	for _, r := range s {
		sMap[r]++
	}

	for _, r := range t {
		tMap[r]++
	}

		if len(sMap) != len(tMap) {
		return false
	}

	for k, v := range sMap {
		val, ok := tMap[k]
		if !ok {
			return false
		}

		if val != v {
			return false
		} 
	}

	return true
}
