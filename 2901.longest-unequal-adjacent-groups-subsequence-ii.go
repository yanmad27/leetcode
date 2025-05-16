package main

/*
 * @lc app=leetcode id=2901 lang=golang
 *
 * [2901] Longest Unequal Adjacent Groups Subsequence II
 */

// @lc code=start

func getHammingDistance(word1 string, word2 string) int {
	if len(word1) != len(word2) {
		return 0
	}
	count := 0
	for i := 0; i < len(word1); i++ {
		if word1[i] != word2[i] {
			count++
		}
	}
	return count
}
func getWordsInLongestSubsequence(words []string, groups []int) []string {
	n := len(words)
	dp := make([]int, n)
	dp[0] = 1
	maxIndex := 0
	maxValue := 1
	for i := 1; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			isHamming := getHammingDistance(words[i], words[j]) == 1
			isDiffGroup := groups[i] != groups[j]
			if isHamming && isDiffGroup {
				dp[i] = max(dp[i], dp[j]+1)
				if dp[i] > maxValue {
					maxValue = dp[i]
					maxIndex = i
				}
			}
		}
	}
	result := []string{words[maxIndex]}
	prevIndex := maxIndex
	for i := maxIndex - 1; i >= 0; i-- {
		if dp[i] != dp[prevIndex]-1 {
			continue
		}
		isHamming := getHammingDistance(words[prevIndex], words[i]) == 1
		isDiffGroup := groups[prevIndex] != groups[i]
		if isHamming && isDiffGroup {
			result = append(result, words[i])
			prevIndex = i
		}
	}

	m := len(result)
	revert := make([]string, m)
	for i := 0; i < m; i++ {
		revert[i] = result[m-i-1]
	}
	return revert
}

// @lc code=end
