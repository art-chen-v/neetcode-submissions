import "slices"

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	res := make([][]int, 0)
	for i := range nums {
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}
		a := i + 1
		b := len(nums) - 1
		for a < b {
			sum := nums[i] + nums[a] + nums[b]
			if sum < 0 {
				a++
			} else if sum > 0 {
				b--
			} else {
				res = append(res, []int{nums[i], nums[a], nums[b]})
				a++
				b--
				for a < b && nums[a] == nums[a - 1] {
					a++
				}
			}
		}
	}
	return res
}
