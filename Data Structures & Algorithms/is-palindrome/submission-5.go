func isPalindrome(s string) bool {
	// "Was it a car or a cat I saw?"
	// No lemon, no melon
	s = strings.ToLower(s)
	l, r := 0, len(s) - 1
	for l <= r {
		if !isAlphaNumeric(s[l]) && l <= r {
			l++
			continue
		}
		if !isAlphaNumeric(s[r]) && l <= r {
			r--
			continue
		}
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func isAlphaNumeric(c byte) bool {
	if (c >= 'A' && c <='Z') || (c >= 'a' && c <='z') || (c >= '0' && c <= '9') {
		return true
	}
	return false
}

