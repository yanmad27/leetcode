package main

import "strings"

/*
 * @lc app=leetcode id=1455 lang=golang
 *
 * [1455] Check If a Word Occurs As a Prefix of Any Word in a Sentence
 */

// @lc code=start
func isPrefixOfWord(sentence string, searchWord string) int {

	words := strings.Split(sentence, " ")
	for i, word := range words {
		if strings.HasPrefix(word, searchWord) {
			return i + 1
		}
	}
	return -1
}

// @lc code=end
