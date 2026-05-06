func maxProfit(prices []int) int {
	maxProfit := 0
	l, r := 0, 1
	for r < len(prices) {
		if prices[l] < prices[r] {
			sum := prices[r] - prices[l]
			maxProfit = max(maxProfit, sum)
		} else {
			l = r
		}
		r++
	}
	return maxProfit
}
