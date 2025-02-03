package main

/*
 * @lc app=leetcode id=567 lang=golang
 *
 * [567] Permutation in String
 */

// @lc code=start

func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}
	for i := 0; i < len(s2)-len(s1)+1; i++ {
		if check1(s1, s2[i:i+len(s1)]) {
			return true
		}
	}
	return false
}

func check1(s1, s2 string) bool {
	s1Arr := make([]int, 26)
	s2Arr := make([]int, 26)
	for _, char := range s1 {
		s1Arr[char-'a']++
	}
	for _, char := range s2 {
		s2Arr[char-'a']++
	}
	for _, char := range s1 {
		if s1Arr[char-'a'] != s2Arr[char-'a'] {
			return false
		}
	}
	return true
}

// @lc code=end
