package problems

/*
 * @lc app=leetcode id=724 lang=golang
 *
 * [724] Find Pivot Index
 */

// @lc code=start
func pivotIndex(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}

	left := 0
	for i, n := range nums {
		// right = total - left - n
		if left == total-left-n {
			return i
		}
		left += n
	}
	return -1
}

// @lc code=end
