func findMin(nums []int) int {
	// 3,4,5,6,1,2
	// 3,1,2
	minNum := nums[0]
	l, r := 0, len(nums) - 1
	for l <= r {
		m := l + (r - l) / 2
		if nums[l] <= nums[m] {
			minNum = min(minNum, nums[l])
			l = m + 1
		} else {
			minNum = min(minNum, nums[m])
			r = m - 1
		}
	}
	return minNum
}
