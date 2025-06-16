package main

import "strconv"

/*
 * @lc app=leetcode id=1432 lang=golang
 *
 * [1432] Max Difference You Can Get From Changing an Integer
 */

// @lc code=start
func maxOperation(numStr string) int {
	arr := []byte(numStr)
	var val byte
	for i, num := range arr {
		if num != '9' {
			val = arr[i]
			break
		}
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] == val {
			arr[i] = '9'
		}
	}
	rs, _ := strconv.Atoi(string(arr))
	return rs
}

func minOperation(numStr string) int {
	arr := []byte(numStr)
	var val byte
	var replacement byte
	if arr[0] != '1' {
		val = arr[0]
		replacement = '1'
	} else {
		founded := false
		for i := 1; i < len(arr); i++ {
			if arr[i] != '1' && arr[i] != '0' {
				val = arr[i]
				replacement = '0'
				founded = true
				break;
			}
		}
		if !founded {
			rs, _ := strconv.Atoi(string(arr))
			return rs
		}
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] == val {
			arr[i] = replacement
		}
	}
	rs, _ := strconv.Atoi(string(arr))
	return rs
}
func maxDiff(num int) int {
	numStr := strconv.Itoa(num)
	return maxOperation(numStr) - minOperation(numStr)
}

// @lc code=end
