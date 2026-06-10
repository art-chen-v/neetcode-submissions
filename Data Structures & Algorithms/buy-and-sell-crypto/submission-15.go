func maxProfit(prices []int) int {
	// [10,1,5,6,7,1]
	res := 0
	l, r := 0, 1
	for r < len(prices) {
		if prices[l] < prices[r] {
			res = max(res, prices[r] - prices[l])
		} else {
			l = r
		}
		r++
	}
	return res
}
