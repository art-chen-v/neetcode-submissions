func maxArea(heights []int) int {
	max := 0

	l, r := 0, len(heights) - 1

	for l < r {
		h := min(heights[l], heights[r])
		w := r - l
		m := h * w
		if m > max {
			max = m
		}
		if heights[l] <= heights[r] {
			l++
		} else {
			r--
		}
	}
	return max
}
