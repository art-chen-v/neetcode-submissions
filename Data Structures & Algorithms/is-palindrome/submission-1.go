func isPalindrome(s string) bool {
	loweredS := strings.ToLower(s)
	l, r := 0, len(loweredS) - 1
	for l < r {
		if !isAlphaNumeric(loweredS[l]) {
			l++
		} else if !isAlphaNumeric(loweredS[r]) {
			r--
		} else {
			if loweredS[l] != loweredS[r] {
				return false
			}
			l++
			r--
		}
	}
	return true
}

func isAlphaNumeric(strByte uint8) bool {
	if (strByte >= 'A' && strByte <= 'Z') ||
		(strByte >= 'a' && strByte <= 'z') ||
		(strByte >= '0' && strByte <= '9') {
			return true
		}
	return false
}