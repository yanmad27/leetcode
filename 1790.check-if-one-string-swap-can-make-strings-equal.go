package main

/*
 * @lc app=leetcode id=1790 lang=golang
 *
 * [1790] Check if One String Swap Can Make Strings Equal
 */

// @lc code=start
func areAlmostEqual(s1 string, s2 string) bool {
	count := 0
	indexs := []int{}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			count++
			indexs = append(indexs, i)
		}
	}
	if count == 2 {
		if s1[indexs[0]] != s2[indexs[1]] || s1[indexs[1]] != s2[indexs[0]] {
			return false
		} else {
			return true
		}
	}
	return count == 0
}

// @lc code=end
