func findMin(nums []int) int {
	res := nums[0]
	if nums[0] < nums[len(nums) - 1] {
		return nums[0]
	}
	l, r := 0, len(nums) - 1
	for l <= r {
		mid := l + (r - l) / 2
		if nums[l] <= nums[mid] {
			res = min(res, nums[l])
			l = mid + 1
		} else {
			r = mid
		}
	}
	return res
}
