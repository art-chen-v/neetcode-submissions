func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for j, a := range nums {
		if i, ok := numsMap[target-a]; ok {
			return []int{i, j} 
		}
		numsMap[a] = j
	}
	return []int{}
}
