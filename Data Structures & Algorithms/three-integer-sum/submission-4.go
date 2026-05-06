import "slices"

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	res := make([][]int, 0)
	for i, a := range nums {
		if i > 0 && a == nums[i-1] {
			continue
		}
		l, r := i + 1, len(nums) - 1
		for l < r {
			sum := a + nums[l] + nums[r]
			if sum < 0 {
				l++
			} else if sum > 0 {
				r--
			} else {
				res = append(res, []int{a, nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}
		}
	}
	return res
}
