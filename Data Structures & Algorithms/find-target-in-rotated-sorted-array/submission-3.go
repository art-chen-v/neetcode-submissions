func search(nums []int, target int) int {
	// 6, 1, 2,   3, 4, 5
	l, r := 0, len(nums) - 1
	for l <= r {
		m := l + (r-l) / 2
		if target == nums[m] {
			return m
		}
		if nums[l] <= nums[m] {
			if target < nums[l] || target > nums[m] {
				l = m + 1
			} else {
				r = m
			}
		} else if nums[m] <= nums[r] {
			if target < nums[m] || target > nums[r] {
				r = m - 1
			} else {
				l = m
			}
		}
	}
	return -1
}
