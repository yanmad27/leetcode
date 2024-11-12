package main

/*
 * @lc app=leetcode id=1829 lang=golang
 *
 * [1829] Maximum XOR for Each Query
 */

// @lc code=start
func flipBit(n, k int) int {
	mask := (1 << k) - 1
	return n ^ mask
}
func getMaximumXor(nums []int, maximumBit int) []int {
	xor := 0
	ans := make([]int, len(nums))
	for i, num := range nums {
		xor ^= num
		ans[len(nums)-i-1] = flipBit(xor, maximumBit)

	}
	return ans
}

// @lc code=end
