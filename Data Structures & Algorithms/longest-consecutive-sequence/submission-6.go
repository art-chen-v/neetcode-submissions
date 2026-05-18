func longestConsecutive(nums []int) int {
	set := make(map[int]struct{})
	for _, n := range nums {
		set[n] = struct{}{}
	}
	longest := 0
	for k, _ := range set {
		if _, ok := set[k-1]; !ok {
			counter := 1
			a := k
			for {
				_, ok := set[a+1]
				if !ok {
					longest = max(longest, counter)
					break
				}
				a += 1
				counter++
			}
		}
	}
	return longest
}
