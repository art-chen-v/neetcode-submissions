func topKFrequent(nums []int, k int) []int {
	bucket := make([][]int, len(nums) + 1)
	group := make(map[int]int)

	for _, v := range nums {
		group[v]++
	}

	for k, v := range group {
		bucket[v] = append(bucket[v], k)
	}

	res := make([]int, 0)
	for i := len(bucket) - 1; i > 0; i-- {
		for _, v := range bucket[i] {
			res = append(res, v)
			if len(res) == k {
				return res
			}
		}
	}
	return res
}
