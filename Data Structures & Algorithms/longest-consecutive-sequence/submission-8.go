func longestConsecutive(nums []int) int {
	longest := 0
	set := make(map[int]struct{})
	for _, n := range nums {
		set[n] = struct{}{}
	}
// 1, 2, 3, 5, 6, 7
	for _, n := range nums {
		if _, ok := set[n-1]; !ok {
			length := 1
			for {
				if _, ok := set[n+1]; !ok {
					longest = max(longest, length)
					break
				}
				n++
				length++
			}
		}
	}
	return longest
}
