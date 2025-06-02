package main

/*
 * @lc app=leetcode id=135 lang=golang
 *
 * [135] Candy
 */

// @lc code=start
func candy(ratings []int) int {

	n := len(ratings)
	candies := make([]int, n)
	for i := range candies {
		candies[i] = 1
	}
	for i := range n - 1 {
		if ratings[i+1] > ratings[i] {
			candies[i+1] = candies[i] + 1
		}
	}
	for i := n - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] &&
			candies[i-1] <= candies[i]+1 {
			candies[i-1] = candies[i] + 1
		}
	}
	sum := 0
	for _, candy := range candies {
		sum += candy
	}
	return sum
}

// @lc code=end
