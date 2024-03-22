package main

import "fmt"

//Find minimum number of coins whose sum make the given amount. Coin values [1,5,10,50,100,500].  Example -  For 12 we need 3 coins, for 468 we need 10 coins.

func findMinimumCoins(coins []int, target int) int {
	count := 0
	for i := len(coins) - 1; i >= 0; i-- {
		for target >= coins[i] {
			target -= coins[i]
			count++
		}
	}
	return count
}

func main() {
	coins := []int{1, 5, 10, 50, 100, 500}
	amount := 12

	fmt.Printf("For %d we need %d coins\n", amount, findMinimumCoins(coins, amount))

	amount = 468
	fmt.Printf("For %d we need %d coins\n", amount, findMinimumCoins(coins, amount))
}
