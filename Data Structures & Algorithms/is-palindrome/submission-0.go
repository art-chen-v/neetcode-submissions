func isPalindrome(s string) bool {
	lowerS := strings.ToLower(s)
	left := 0
	right := len(lowerS) - 1
	for left < right {
		if !checkAlphaNum(lowerS[left]) {
			left++
			continue
		}
		if !checkAlphaNum(lowerS[right]) {
			right--
			continue
		}
		if lowerS[left] != lowerS[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func checkAlphaNum(strByte uint8) bool {
	if (strByte >= '0' && strByte <= '9') ||
    	(strByte >= 'A' && strByte <= 'Z') ||
		(strByte >= 'a' && strByte <= 'z') {
			return true
		}
	return false
}
