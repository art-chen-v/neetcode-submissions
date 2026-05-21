func maxArea(heights []int) int {
	res := 0
	l, r := 0, len(heights) - 1
	for l < r {
		size := (r-l) * min(heights[l], heights[r])
		res = max(res, size)
		if heights[l] < heights[r] {
			l++
		} else{
			r--
		}
	}
	return res
}
