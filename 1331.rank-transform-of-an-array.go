/*
 * @lc app=leetcode id=1331 lang=golang
 *
 * [1331] Rank Transform of an Array
 */

// @lc code=start
func arrayRankTransform(arr []int) []int {
	clone := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		clone[i] = arr[i]
	}
	sort.Ints(clone)
	m := make(map[int]int)
	rank := 1
	for i := 0; i < len(clone); i++ {
		if _, ok := m[clone[i]]; !ok {
			m[clone[i]] = rank
			rank++
		}
	}
	for i := 0; i < len(arr); i++ {
		arr[i] = m[arr[i]]
	}
	return arr
}

// @lc code=end

