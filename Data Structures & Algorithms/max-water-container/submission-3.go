func maxArea(heights []int) int {
	// [1,7,2,5,4,7,3,6]
	// [1,2,1]
	l, r := 0, len(heights) - 1
	maxArea := 0
	for l < r {
		w := r - l
		h := min(heights[l], heights[r])
		maxArea = max(maxArea, w * h)
		if heights[l] < heights[r] {
			l++
		} else {
			r--
		}
	}

	return maxArea
}
