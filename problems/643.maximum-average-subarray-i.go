package problems

/*
 * @lc app=leetcode id=643 lang=golang
 *
 * [643] Maximum Average Subarray I
 */

// @lc code=start
func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	average := float64(sum) / float64(k)
	for i := k; i < len(nums); i++ {
		sum += nums[i]
		sum -= nums[i-k]
		average = max(average, float64(sum)/float64(k))
	}
	return average
}

// @lc code=end
