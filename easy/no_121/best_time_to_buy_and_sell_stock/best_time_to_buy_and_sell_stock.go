package best_time_to_buy_and_sell_stock

func maxProfitClearer(prices []int) int {
	preRecord := 0
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		preRecord = GetMax(preRecord-prices[i-1]+prices[i], 0)
		if preRecord > maxProfit {
			maxProfit = preRecord
		}
	}
	return maxProfit
}

func GetMax(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxProfitBetter(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	min := prices[0]
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}
		profit := prices[i] - min
		if profit > maxProfit {
			maxProfit = profit
		}
	}
	return maxProfit
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	max := 0
	minMap := updateMinMap(prices)
	maxMap := updateMaxMap(prices)

	for i := 0; i < len(prices)-1; i++ {
		result := maxMap[i+1] - minMap[i]
		if result > max {
			max = result
		}
	}
	return max
}

func updateMaxMap(prices []int) map[int]int {
	result := make(map[int]int, len(prices))
	result[len(prices)-1] = prices[len(prices)-1]
	for i := len(prices) - 2; i >= 0; i-- {
		if prices[i] > result[i+1] {
			result[i] = prices[i]
		} else {
			result[i] = result[i+1]
		}
	}

	return result
}

func updateMinMap(prices []int) map[int]int {
	result := make(map[int]int, len(prices))
	result[0] = prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < result[i-1] {
			result[i] = prices[i]
		} else {
			result[i] = result[i-1]
		}
	}

	return result
}
