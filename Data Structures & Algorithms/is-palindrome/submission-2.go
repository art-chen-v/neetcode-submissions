func isPalindrome(s string) bool {
	// "tab a cat"
	loweredS := strings.ToLower(s)
	l, r := 0, len(loweredS) - 1
	for l <= r {
		if (loweredS[l] < 48 || loweredS[l] > 57) && (loweredS[l] < 97 || loweredS[l] > 122) {
			l++
			continue
		}

		if (loweredS[r] < 48 || loweredS[r] > 57) && (loweredS[r] < 97 || loweredS[r] > 122) {
			r--
			continue
		}

		if loweredS[l] != loweredS[r] {
			return false
		}

		l++
		r--
	}
	return true
}
