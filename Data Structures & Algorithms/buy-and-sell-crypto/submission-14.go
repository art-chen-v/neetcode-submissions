func maxProfit(prices []int) int {
// [2,1,2,1,0,1,2]
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
