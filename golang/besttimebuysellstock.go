/* Best Time to Buy and Sell Stock II

Input: prices = [7,1,5,3,6,4]
Output: 7
Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.

prices = [7,1,5,3,6,4]

Day-to-day differences:
7 → 1 = -6 (ignore, price going down)
1 → 5 = +4 (profit! take it)
5 → 3 = -2 (ignore)
3 → 6 = +3 (profit! take it)
6 → 4 = -2 (ignore)
Total profit = 4 + 3 = 7 */

package main

import "fmt"

func main() {
  prices := []int{7,1,5,3,6,4}
  fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	var res int
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}
	}
	return res
}
