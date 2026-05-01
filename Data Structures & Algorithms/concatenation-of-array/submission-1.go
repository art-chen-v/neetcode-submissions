func getConcatenation(nums []int) []int {
    ans := make([]int, len(nums)*2)
	for i, v := range nums {
		ans[i] = v
		ans[i+len(nums)] = v
	}
	return ans
}
