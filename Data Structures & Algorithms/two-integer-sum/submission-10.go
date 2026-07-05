func twoSum(nums []int, target int) []int {
	// [1,3,4,2]
	// 6
	m := make(map[int]int)
	// res := make([]int, 2)
	for i, n := range nums {
		m[n] = i
	}
	for i, n := range nums {
		if index, ok := m[target - n]; ok && i != index {
			return []int{i, index}
		}
	}
	return []int{}
}
