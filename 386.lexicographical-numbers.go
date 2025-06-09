package main

/*
 * @lc app=leetcode id=386 lang=golang
 *
 * [386] Lexicographical Numbers
 */

// @lc code=start
func lexicalOrder(n int) []int {

	result := make([]int, n)
	cur := 1
	for i := 0; i < n; i++ {
		result[i] = cur
		if cur*10 <= n {
			cur *= 10
		} else {
			if cur >= n {
				cur /= 10
			}
			cur++
			for cur%10 == 0 {
				cur /= 10
			}
		}
	}
	return result
}

// @lc code=end
