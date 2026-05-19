func findMin(nums []int) int {
	l, r := 0, len(nums) - 1
	res := nums[0]
// 5, 0
	for l <= r {
		mid := l + (r-l) / 2
		if nums[l] <= nums[mid] {
			res = min(res, nums[l])
			l = mid + 1
		} else {
			r = mid
		}
	}

	return res
}
