func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	ans, max := 0, prices[len(prices)-1]
	for i := len(prices) - 2; i >= 0; i-- {
		if max-prices[i] > ans {
			ans = max - prices[i]
		}
		if prices[i] > max {
			max = prices[i]
		}
	}
	return ans
}
