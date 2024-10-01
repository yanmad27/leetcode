package main

/*
 * @lc app=leetcode id=1497 lang=golang
 *
 * [1497] Check If Array Pairs Are Divisible by k
 */

// @lc code=start
func canArrange(arr []int, k int) bool {

	rest := make([]int, k)
	for _, num := range arr {
		num = num % k
		if num < 0 {
			num += k
		}
		rest[num]++
	}
	if rest[0]%2 != 0 {
		return false
	}
	for i := 1; i < k; i++ {
		if rest[i] != rest[k-i] {
			return false
		}
	}
	return true
}

// @lc code=end
