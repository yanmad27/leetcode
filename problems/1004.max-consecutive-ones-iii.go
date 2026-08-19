package problems

/*
 * @lc app=leetcode id=1004 lang=golang
 *
 * [1004] Max Consecutive Ones III
 */

// @lc code=start
func longestOnes(nums []int, k int) int {

	n := len(nums)
	left, right := 0, 0
	countZero := 0
	result := 0

	for right < n {
		if nums[right] == 0 {
			countZero++
		}
		if countZero > k {
			for countZero > k {
				if nums[left] == 0 {
					countZero--
				}
				left++
			}
		} else {
			result = max(result, right-left+1)
		}
		right++
	}
	return result
}

// @lc code=end
