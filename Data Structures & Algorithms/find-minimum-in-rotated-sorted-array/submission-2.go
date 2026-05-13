func findMin(nums []int) int {
	l, r := 0, len(nums) - 1
	minV := math.MaxInt
	if nums[l] <= nums[r] {
		return nums[l]
	}
	for l <= r {
		mid := (r + l) / 2
		left := nums[l:mid + 1]
		right := nums[mid + 1:]
		if left[0] <= left[len(left) - 1] && right[0] <= right[len(right) - 1] {
			return min(left[0], right[0])
		} else {
			if left[0] <= left[len(left) - 1] {
				minV = min(minV, left[0])
			} else {
				r = mid
			}
			if right[0] <= right[len(right) - 1] {
				minV = min(minV, right[0])
			} else {
				l = mid + 1
			}
		}
	}
	return minV
}
