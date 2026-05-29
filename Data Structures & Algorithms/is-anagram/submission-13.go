func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sMap := make(map[byte]int)
	tMap := make(map[byte]int)
	for i := range s {
		sMap[s[i]]++
		tMap[t[i]]++
	}
	for i := range sMap {
		if sMap[i] != tMap[i] {
			return false
		}
	}
	return true
}
