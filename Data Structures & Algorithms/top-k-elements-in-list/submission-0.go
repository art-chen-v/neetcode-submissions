func topKFrequent(nums []int, k int) []int {
	counter := make(map[int]int)
	bucket := make([][]int, len(nums)+1)

	for _, v := range nums {
		counter[v]++
	}

	for k, v := range counter {
		bucket[v] = append(bucket[v], k)
	}

	res := make([]int, 0)

	for i := len(bucket) - 1; i >= 1; i-- {
		for _, n := range bucket[i] {
			res = append(res, n)
			if len(res) == k {
				return res
			}
		}
	}
	return res
}
