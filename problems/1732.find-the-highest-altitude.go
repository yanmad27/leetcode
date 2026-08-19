package problems

/*
 * @lc app=leetcode id=1732 lang=golang
 *
 * [1732] Find the Highest Altitude
 */

// @lc code=start
func largestAltitude(gain []int) int {
	rs := 0
	cur := 0
	for _, v := range gain {
		cur += v
		rs = max(rs, cur)
	}
	return rs
}

// @lc code=end
