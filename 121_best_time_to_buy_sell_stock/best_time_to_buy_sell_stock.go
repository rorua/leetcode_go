package best_time_to_buy_sell_stock

func maxProfit(prices []int) int {
	var profit int = 0
	l, r := 0, 1
	for r < len(prices) {
		if prices[l] < prices[r] {
			if prices[r]-prices[l] > profit {
				profit = prices[r] - prices[l]
			}
		} else {
			l = r
		}
		r++
	}
	return profit
}
