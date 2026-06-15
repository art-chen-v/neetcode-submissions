func search(nums []int, target int) int {
	// nums=[3,5,6,0,1,2]
	// target=4
	l, r := 0, len(nums) - 1
	for l <= r {
		m := l + (r - l) / 2
		if nums[m] == target {
			return m
		}
		if nums[l] <= nums[m] {
			if nums[l] <= target && nums[m] >= target {
				r = m
			} else {
				l = m + 1
			}
		} else if nums[m] < nums[r] {
			if nums[m] <= target && nums[r] >= target {
				l = m
			} else {
				r = m -1
			}
		}
	}

	return -1
}
