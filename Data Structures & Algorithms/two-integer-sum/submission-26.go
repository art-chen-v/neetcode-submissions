func twoSum(nums []int, target int) []int {
    visited := make(map[int]int)
	for i, n := range nums {
		if index, ok := visited[target - n]; ok {
			return []int{index, i}
		}
		visited[n] = i
	}
	return []int{}
}
