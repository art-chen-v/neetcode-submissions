func maxArea(heights []int) int {
	maxWater := 0

	l, r := 0, len(heights) - 1

	for l < r {
		w := r - l
		h := min(heights[l], heights[r])
		result := w * h
		maxWater = max(maxWater, result)
		if heights[l] <= heights[r] {
			l++
		} else {
			r--
		}
	}
	return maxWater
}
