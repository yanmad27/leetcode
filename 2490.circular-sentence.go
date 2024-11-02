package main

/*
 * @lc app=leetcode id=2490 lang=golang
 *
 * [2490] Circular Sentence
 */

// @lc code=start
func isCircularSentence(sentence string) bool {

	for i := 0; i < len(sentence)-1; i++ {
		if sentence[i] == ' ' && sentence[i-1] != sentence[i+1] {
			return false
		}
	}
	return sentence[0] == sentence[len(sentence)-1]
}

// @lc code=end
