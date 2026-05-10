func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	loweredS := strings.ToLower(s)
	loweredT := strings.ToLower(t)

	sCounter := [26]int{}
	tCounter := [26]int{}

	for i := 0; i < len(loweredS); i++ {
		sCounter[loweredS[i]-'a']++
		tCounter[loweredT[i]-'a']++
	}
	return sCounter == tCounter
}
