package main

import (
	"sort"
)

/*
 * @lc app=leetcode id=3169 lang=golang
 *
 * [3169] Count Days Without Meetings
 */

// @lc code=start
func countDays(days int, meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool {
		if meetings[i][0] < meetings[j][0] {
			return true
		}
		if meetings[i][0] == meetings[j][0] {
			return meetings[i][1] < meetings[j][1]
		}
		return false
	})
	mergedMeetings := make([][]int, 0)
	for i := 0; i < len(meetings); i++ {
		newMeet := meetings[i]
		for j := i + 1; j < len(meetings); j++ {
			meet := meetings[j]
			if newMeet[1] < meet[0]-1 {
				break
			}
			if newMeet[1] >= meet[1] {
				i++
				continue
			}
			if newMeet[1] <= meet[1] {
				newMeet[1] = meet[1]
				i++
			}
		}
		mergedMeetings = append(mergedMeetings, newMeet)
	}

	result := days
	for _, meet := range mergedMeetings {
		result -= meet[1] - meet[0] + 1
	}
	return result
}

// @lc code=end
