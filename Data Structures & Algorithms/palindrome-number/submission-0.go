func isPalindrome(x int) bool {
	num := x
	res := 0
	for num > 0 {
		res = res * 10 + num % 10
		num = num / 10
	}
	return res == x
}
