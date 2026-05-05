func twoSum(nums []int, target int) []int {
    prevNumsMap := make(map[int]int)

	for i, n := range nums {
		index, ok := prevNumsMap[target-n]
		if ok {
			return []int{index, i}
		}
		prevNumsMap[n] = i
	}

	return []int{}
}
