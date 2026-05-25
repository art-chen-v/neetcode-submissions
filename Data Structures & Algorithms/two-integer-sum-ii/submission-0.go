func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers) - 1
	for l < r {
		if l > 0 && numbers[l] == numbers[l-1] {
			continue
		}
		sum := numbers[l] + numbers[r]
		if sum < target {
			l++
		} else if sum > target {
			r--
		} else {
			return []int{l+1, r+1}
		}
	}
	return []int{}
}
