package main

func maxProfit(prices []int) int {
	count := len(prices)
	if count < 2 {
		return 0
	}

	maxProfit := 0
	for i := 1; i < count; i++ {
		if prices[i] > prices[i-1] {
			maxProfit += prices[i] - prices[i-1]
		}
	}

	return maxProfit
}
