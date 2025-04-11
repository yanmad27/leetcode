package main

/*
 * @lc app=leetcode id=2843 lang=golang
 *
 * [2843]   Count Symmetric Integers
 */

// @lc code=start
func isSymmetric(num int) bool {
	if num < 10 {
		return false
	}
	if num < 100 {
		return num%10 == num/10
	}
	if num < 1000 {
		return false
	}
	if num < 10000 {
		return num%10+(num/10%10) == (num/100%10 + num/1000)
	}

	return true
}
func countSymmetricIntegers(low int, high int) int {
	count := 0
	for i := low; i <= high; i++ {
		if isSymmetric(i) {
			count++
		}
	}
	return count
}

// @lc code=end
