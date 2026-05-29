func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var counterS [26]int
	var counterT [26]int
	for i := range s {
		counterS[s[i] - 'a']++
		counterT[t[i] - 'a']++
	}
	return counterS == counterT
}
