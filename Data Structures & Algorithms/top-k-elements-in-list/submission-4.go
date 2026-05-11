func topKFrequent(nums []int, k int) []int {
	bucketFreq := make([][]int, len(nums)+1)
	counter := make(map[int]int)
	result := make([]int, 0)
	for _, n := range nums {
		counter[n]++
	}

	for k, v := range counter {
		bucketFreq[v] = append(bucketFreq[v], k)
	}

	for i := len(bucketFreq)-1; i > 0; i-- {
		for _, n := range bucketFreq[i] {
			if len(result) < k {
				result = append(result, n)
			} else {
				return result
			}
		}
	}
	return result
}
