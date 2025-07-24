package main

import (
	"fmt"
)

func MaxProfit(prices []int) (buyDay, sellDay, maxProfit int) {
	if len(prices) < 2 {
		return -1, -1, 0 // Not enough days to make a transaction
	}

	minPrice := prices[0]
	minDay := 0
	maxProfit = 0

	for i := 1; i < len(prices); i++ {
		profit := prices[i] - minPrice
		if profit > maxProfit {
			maxProfit = profit
			buyDay = minDay
			sellDay = i
		}
		if prices[i] < minPrice {
			minPrice = prices[i]
			minDay = i
		}
	}

	if maxProfit == 0 {
		return -1, -1, 0 // No profit possible
	}
	return buyDay, sellDay, maxProfit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	buy, sell, profit := MaxProfit(prices)
	if buy == -1 {
		fmt.Println("No profitable transaction possible.")
	} else {
		fmt.Printf("Buy on day %d and sell on day %d for a profit of %d.\n", buy, sell, profit)
	}
}
