func longestConsecutive(nums []int) int {
	set := make(map[int]struct{})
	for _, n := range nums {
		set[n] = struct{}{}
	}
	longest := 0
	for i, _ := range set {
		if _, ok := set[i-1]; !ok {
			sum := 1
			next := i + 1
			for {
				if _, ok := set[next]; !ok {
					longest = max(longest, sum)
					break
				}
				sum++
				next++
			}
		} 
	}
	return longest
}
