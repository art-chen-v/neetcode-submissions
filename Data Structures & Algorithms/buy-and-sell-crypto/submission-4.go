func maxProfit(prices []int) int {
	// [7,1,5,3,6,4]
	l, r := 0, 1
	maxProfit := 0

	for l < r && r < len(prices) {
		if prices[l] <= prices[r] {
			profit := prices[r] - prices[l]
			maxProfit = max(profit, maxProfit)
			r++
		} else {
			l = r
			r++
		}
	}

	return maxProfit
}
