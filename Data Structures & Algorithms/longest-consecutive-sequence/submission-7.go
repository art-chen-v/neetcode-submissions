func longestConsecutive(nums []int) int {
	count := 0
	set := make(map[int]struct{})
	for _, n := range nums {
		set[n] = struct{}{}
	}
// 1, 2, 3, 4, 8, 9, 10
	for k, _ := range set {
		if _, ok := set[k-1]; !ok {
			c := 1
			for {
				if _, ok := set[k+c]; !ok {
					break
				}
				c++
			}
			count = max(count, c)
		}
	}
	return count
}
