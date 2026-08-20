func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	bucket := make([][]int, len(nums) + 1)

	for _, n := range nums {
		freq[n]++
	}

	for k, v := range freq {
		bucket[v] = append(bucket[v], k)
	}

	res := make([]int, 0)
	for i := len(bucket) - 1; i > 0; i-- {
		for _, n := range bucket[i] {
			if k > 0 {
				res = append(res, n)
				k--
			}
		}
	}
	return res
}
