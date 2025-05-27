package main

/*
 * @lc app=leetcode id=2894 lang=golang
 *
 * [2894] Divisible and Non-divisible Sums Difference
 */

// @lc code=start
func differenceOfSums(n int, m int) int {
	sum := n * (n + 1) / 2
	num2 := 0
	tmp := m
	for tmp <= n {
		num2 += tmp
		tmp += m
	}
	num1 := sum - num2
	return num1 - num2
}

// @lc code=end
