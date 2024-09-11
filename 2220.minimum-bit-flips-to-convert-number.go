package leetcode

import "math/bits"

/*
 * @lc app=leetcode id=2220 lang=golang
 *
 * [2220] Minimum Bit Flips to Convert Number
 */

// @lc code=start
func minBitFlips(start int, goal int) int {
	return bits.OnesCount(uint(start ^ goal))
}

// @lc code=end
