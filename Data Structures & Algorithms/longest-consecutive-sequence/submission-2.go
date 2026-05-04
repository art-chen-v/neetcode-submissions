func longestConsecutive(nums []int) int {
	m := make(map[int]struct{})

	for _, n := range nums {
		m[n] = struct{}{}
	}

	max := 0

	for k, _ := range m {
		_, ok := m[k-1]
		if !ok {
			l := 1
			n := k
			for {
				if _, ok := m[n+1]; ok {
					l++
					n += 1
				} else {
					if l > max {
						max = l
					}
					break
				} 
			}
		}
	}

	return max
}
