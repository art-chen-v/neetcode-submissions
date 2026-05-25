func findMin(nums []int) int {
	if nums[0] < nums[len(nums)-1] {
		return nums[0]
	}
	res := nums[0]
	l, r := 0, len(nums) - 1
	for l <= r {
		m := l + (r-l) / 2
		if nums[l] <= nums[m] {
			res = min(res, nums[l])
			l = m + 1
		} else if nums[l] > nums[m] {
			r = m
		}
	}
	return res
}
