package main

/*
 * @lc app=leetcode id=1475 lang=golang
 *
 * [1475] Final Prices With a Special Discount in a Shop
 */

// @lc code=start
func finalPrices(prices []int) []int {
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[i] >= prices[j] {
				prices[i] -= prices[j]
				break
			}
		}
	}
	return prices
}

// @lc code=end
