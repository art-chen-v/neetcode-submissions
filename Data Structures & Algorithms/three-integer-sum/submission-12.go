import "slices"

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	slices.Sort(nums)
	for i := range nums {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}
		j := i + 1
		k := len(nums) - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum < 0 {
				j++ 
			} else if sum > 0 {
				k--
			} else {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for j < k && nums[j] == nums[j - 1] {
					j++
				}
			}
		}
	}
	return res
}
