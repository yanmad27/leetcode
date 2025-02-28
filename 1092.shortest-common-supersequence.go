package main

/*
 * @lc app=leetcode id=1092 lang=golang
 *
 * [1092] Shortest Common Supersequence
 */

// @lc code=start
func shortestCommonSupersequence(str1 string, str2 string) string {
	m, n := len(str1), len(str2)
	result := ""
	maxK := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if str1[i] == str2[j] {
				k := 1
				for ; i+k < m && j+k < n && str1[i+k] == str2[j+k]; k++ {

				}
				if k >= maxK &&
					(i == 0 || i+k == m) &&
					(j == 0 || j+k == n) {
					tmpRs := ""
					if i == 0 {
						tmpRs = str2 + str1[k:]
					} else if j == 0 {
						tmpRs = str1 + str2[k:]
					}
					maxK = k
					if result == "" || len(tmpRs) < len(result) {
						result = tmpRs
					}
				}
			}
		}
	}

	return result
}

// @lc code=end
