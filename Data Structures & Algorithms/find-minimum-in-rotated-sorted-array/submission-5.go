func findMin(nums []int) int {
	res := nums[0]
	l, r := 0, len(nums) - 1
	if nums[l] < nums[r] {
		return nums[l]
	}
	for l <= r {
		m := l + (r - l) / 2
		if nums[l] <= nums[m] {
			res = min(res, nums[l])
			l = m + 1
		} else {
			r = m
		}
	}
	return res
}
