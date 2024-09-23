package main

/*
 * @lc app=leetcode id=2707 lang=golang
 *
 * [2707] Extra Characters in a String
 */
// @lc code=start

func findAllIndex(s, word string) []int {
	res := []int{}
	for i := 0; i < len(s)-len(word)+1; i++ {
		if s[i:i+len(word)] == word {
			res = append(res, i)
		}
	}
	return res
}

func minExtraChar(s string, dictionary []string) int {
	existedMap := make([]bool, len(s))
	for _, word := range dictionary {
		allIndex := findAllIndex(s, word)
		for _, index := range allIndex {
			for i := 0; i < len(word); i++ {
				existedMap[index+i] = true
			}
		}
	}
	count := 0
	for i := 0; i < len(existedMap); i++ {
		if !existedMap[i] {
			count++
		}
	}
	return count
}

// @lc code=end
