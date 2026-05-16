func twoSum(nums []int, target int) []int {
    numsMap := make(map[int]int)
	for i := range nums {
		if index, ok := numsMap[target-nums[i]]; ok {
			return []int{index, i}
		}
		numsMap[nums[i]] = i
	}
	return []int{}
}
