import "slices"

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	// [-4,-1,-1,0,1,2,2]
	// 
	// [-1,-1,2], [-1,0,1]
	res := make([][]int, 0)
	for i, a := range nums {
		if i > 0 && a == nums[i-1] {
			continue
		}
		l := i + 1
		r := len(nums) - 1
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
