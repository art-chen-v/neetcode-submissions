func topKFrequent(nums []int, k int) []int {
	bucket := make([][]int, len(nums)+1)
	counter := make(map[int]int)

	for _, n := range nums {
		counter[n]++
	}

	for k, v := range counter {
		bucket[v] = append(bucket[v], k)
	}

	output := make([]int, 0)

	for i := len(bucket)-1; i > 0; i-- {
		for _, n := range bucket[i] {
			output = append(output, n)
			if len(output) == k {
				return output
			}
		}
	}

	return output
}
