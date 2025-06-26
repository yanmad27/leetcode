package main

/*
 * @lc app=leetcode id=1071 lang=golang
 *
 * [1071] Greatest Common Divisor of Strings
 */

// @lc code=start
func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	s1, s2 := len(str1), len(str2)
	n := min(s1, s2)
	for i := n; i > 0; i-- {
		if str1[:i] == str2[:i] &&
			s1%i == 0 &&
			s2%i == 0 {
			return str1[:i]
		}
	}
	return ""
}

// @lc code=end
