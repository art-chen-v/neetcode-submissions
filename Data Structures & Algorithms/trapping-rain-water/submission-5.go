func trap(height []int) int {
	res := 0
	l, r := 0, 2
	maxWall := [2]int{}
	for r < len(height) {
		if height[l] == 0 {
			l++
			r++
			continue
		}

		maxWall = findMax(maxWall, [2]int{r, height[r]})

		for r < len(height) {
			if height[l] > height[r] {
				r++
			} else {
				break
			}
			if r < len(height) {
				maxWall = findMax(maxWall, [2]int{r, height[r]})
			}
		}

		r = maxWall[0]
		maxWall = [2]int{}
		h := min(height[l], height[r])
		for l+1 < r && height[l+1] >= h {
			l++
			h = min(height[l], height[r])
		}
		inner := l + 1
		sum := (r - l - 1) * h
		for inner < r {
			sum -= height[inner]
			inner++
		}
		if sum > 0 {
			res += sum
		}
		l = r
		r += 2
	}
	return res
}

func findMax(m1, m2 [2]int) [2]int {
	if m1[1] < m2[1] {
		return m2
	}
	return m1
}
