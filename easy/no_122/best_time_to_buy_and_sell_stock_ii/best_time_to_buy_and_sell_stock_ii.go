package best_time_to_buy_and_sell_stock_ii

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	result := 0
	min := -1
	if prices[1] >= prices[0] {
		min = 0
	}
	for i := 1; i < len(prices)-1; i++ {
		if min == -1 {
			if prices[i] <= prices[i-1] && prices[i] <= prices[i+1] {
				min = i
			}
		} else {
			if prices[i] >= prices[i-1] && prices[i] >= prices[i+1] {
				result = result + prices[i] - prices[min]
				min = -1
			}
		}
	}
	if prices[len(prices)-1] >= prices[len(prices)-2] && min != -1 {
		result = result + prices[len(prices)-1] - prices[min]
	}

	return result
}


func maxProfitBetter(prices []int) int{
	result := 0
	for i := 0; i < len(prices) - 1 ; i++ {
		if prices[i] < prices[i + 1] {
			result += prices[i + 1] - prices[i]
		}
	}
	return result
}