func topKFrequent(nums []int, k int) []int {
	bucket := make([][]int, len(nums) + 1)
	freqMap := make(map[int]int)
	for _, v := range nums {
		freqMap[v]++
	}
	for n, v := range freqMap {
		bucket[v] = append(bucket[v], n)
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
