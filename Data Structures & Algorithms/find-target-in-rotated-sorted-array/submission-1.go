func search(nums []int, target int) int {
	l, r := 0, len(nums) - 1
	for l <= r {
		mid := l + (r - l) / 2

		if target == nums[mid] {
			return mid
		}

		if nums[l] <= nums[mid] {
			if target >= nums[l] && target <= nums[mid] {
				r = mid
			} else {
				l = mid + 1
			}
		} else if nums[mid] <= nums[r] {
			if target >= nums[mid] && target <= nums[r] {
				l = mid
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}
