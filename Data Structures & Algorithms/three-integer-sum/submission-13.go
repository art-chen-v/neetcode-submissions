import "slices"

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	res := make([][]int, 0)
	// [-4, -1,-1,-1,0,1,2]
	for i := range nums {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		j := i + 1
		k := len(nums) - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum > 0 {
				k--
			} else if sum < 0 {
				j++ 
			} else {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for j < k && nums[j] == nums[j-1] {
					j++
				}
			}
		}
	}
	return res
}
