package main

/*
 * @lc app=leetcode id=2559 lang=golang
 *
 * [2559] Count Vowel Strings in Ranges
 */

// @lc code=start
func vowelStrings(words []string, queries [][]int) []int {
	vowelMap := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}

	prefix := make([]int, len(words)+1)
	for i := 0; i < len(words); i++ {
		if vowelMap[words[i][0]] &&
			vowelMap[words[i][len(words[i])-1]] {
			prefix[i+1] = prefix[i] + 1
		} else {
			prefix[i+1] = prefix[i]
		}
	}
	rs := make([]int, len(queries))
	for i, query := range queries {
		rs[i] = prefix[query[1]+1] - prefix[query[0]]
	}

	return rs
}

// @lc code=end
