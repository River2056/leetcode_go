package best_time_to_buy_and_sell_stock

func MaxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	profit := 0
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else {
			if prices[i] - min > profit {
				profit = prices[i] - min
			}
		}
	}
	return profit
}
