func longestConsecutive(nums []int) int {
	numMap := make(map[int]struct{})
	for _, n := range nums {
		numMap[n] = struct{}{}
	}

	max := 0

	for _, n := range nums {
		longest := 0
		next := n
		if _, ok := numMap[next]; ok {
			longest += 1
			if _, ok := numMap[next-1]; !ok {
				_, ok := numMap[next+1]
				for ok {
					longest += 1
					next++
					_, ok = numMap[next+1]
				}
				if longest > max {
					max = longest
				}
			}
		}
	}
	return max
}
