func hasDuplicate(nums []int) bool {
    set := make(map[int]struct{})
	for _, n := range nums {
		set[n] = struct{}{}
	}
	return len(set) != len(nums)
}
