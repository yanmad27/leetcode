package main

/*
 * @lc app=leetcode id=1456 lang=golang
 *
 * [1456] Maximum Number of Vowels in a Substring of Given Length
 */

// @lc code=start
func maxVowels(s string, k int) int {
	vowels := make([]bool, 'z'+1)
	vowels['a'-'a'] = true
	vowels['i'-'a'] = true
	vowels['u'-'a'] = true
	vowels['e'-'a'] = true
	vowels['o'-'a'] = true

	count := 0
	for i := range k {
		if vowels[s[i]-'a'] {
			count++
		}
	}
	result := count
	for i := k; i < len(s); i++ {
		if vowels[s[i]-'a'] {
			count++
		}
		if vowels[s[i-k]-'a'] {
			count--
		}
		result = max(result, count)
	}
	return result
}

// @lc code=end
