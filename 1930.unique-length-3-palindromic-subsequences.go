package main

/*
 * @lc app=leetcode id=1930 lang=golang
 *
 * [1930] Unique Length-3 Palindromic Subsequences
 */

// @lc code=start
func countPalindromicSubsequence(s string) int {

	n := len(s)
	count := 0
	for i := 'a'; i <= 'z'; i++ {
		for j := 'a'; j <= 'z'; j++ {
			for k := 0; k < n-2; k++ {
				if rune(s[k]) == i {
					for h := k + 1; h < n-1; h++ {
						if rune(s[h]) == j {
							for m := h + 1; m < n; m++ {
								if rune(s[m]) == i {
									count++
									break
								}
							}
							break
						}
					}
					break
				}
			}
		}
	}

	return count
}

// @lc code=end
