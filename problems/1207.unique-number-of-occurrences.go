package problems

/*
 * @lc app=leetcode id=1207 lang=golang
 *
 * [1207] Unique Number of Occurrences
 */

// @lc code=start
func uniqueOccurrences(arr []int) bool {
	countMaps := make(map[int]int)
	for _, v := range arr {
		countMaps[v]++
	}
	countMaps2 := make(map[int]bool)
	for _, val := range countMaps {
		if !countMaps2[val] {
			countMaps2[val] = true
		} else {
			return false
		}
	}
	return true
}

// @lc code=end
