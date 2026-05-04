func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
	res := make([]int, 0)
	for i, v := range nums {
		if index, ok := m[target-v]; ok {
			res = append(res, index)
			res = append(res, i)
			break
		} else {
			m[v] = i
		}
	}
	return res
}
