func search(nums []int, target int) int {
	// 1, 2, 3, 4, 5, 6, 7, 8
	// 0  1  2  3  4  5  6  7
	l, r := 0, len(nums) - 1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
