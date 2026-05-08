func maxProfit(prices []int) int {
	mProfit := 0
	l, r := 0, 1

	for r < len(prices) {
		if prices[l] > prices[r] {
			l++
		} else {
			mProfit = max(mProfit, prices[r] - prices[l])
			r++
		}
	}
	return mProfit
}
