func topKFrequent(nums []int, k int) []int {
	bucket := make([][]int, len(nums) + 1)
	counter := make(map[int]int)
	for _, n := range nums {
		counter[n]++
	}
	for k, v := range counter {
		bucket[v] = append(bucket[v], k)
	}

	res := make([]int, 0)
	for i := len(bucket) - 1; i > 0; i-- {
		for _, n := range bucket[i] {
			if len(res) < k {
				res = append(res, n)
			}
		}
	}
	return res
}
