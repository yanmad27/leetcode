package main

/*
 * @lc app=leetcode id=2028 lang=golang
 *
 * [2028] Find Missing Observations
 */

// @lc code=start
func missingRolls(rolls []int, mean int, n int) []int {
	sum := mean * (len(rolls) + n)
	for _, roll := range rolls {
		sum -= roll
	}
	if sum < n || sum > 6*n {
		return []int{}
	}
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = 1
	}
	for i := 0; i < sum-n; i++ {
		res[i%n]++
	}
	return res
}

// @lc code=end
