func longestConsecutive(nums []int) int {
	longest := 0
	set := make(map[int]struct{})
	for _, n := range nums {
		set[n] = struct{}{}
	}

	for _, n := range nums {
		if _, ok := set[n-1]; !ok {
			length := 1
			a := n
			for {
				if _, ok := set[a+1]; !ok {
					longest = max(longest, length)
					break
				} else {
					length++
					a++
				}
			}
		} else {
			continue
		}
	}

	return longest
}
