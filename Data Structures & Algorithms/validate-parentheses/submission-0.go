func isValid(s string) bool {
	closeToOpen := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := make([]rune, 0)

	for _, r := range s {
		if v, ok := closeToOpen[r]; ok {
			if len(stack) > 0 && v == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, r)
		}
	}

	return len(stack) == 0
}
