package main

/*
 * @lc app=leetcode id=2381 lang=golang
 *
 * [2381] Shifting Letters II
 */

// @lc code=start

func shiftingLetters(s string, shifts [][]int) string {
	n := len(s)
	shiftArr := make([]int, n+1)

	for _, shift := range shifts {
		start, end, direction := shift[0], shift[1], shift[2]
		if direction == 1 {
			shiftArr[start]++
			shiftArr[end+1]--
		} else {
			shiftArr[start]--
			shiftArr[end+1]++
		}
	}

	currentShift := 0
	result := make([]byte, n)
	for i := 0; i < n; i++ {
		currentShift += shiftArr[i]
		shifted := (int(s[i]-'a') + currentShift) % 26
		if shifted < 0 {
			shifted += 26
		}
		result[i] = byte(shifted) + 'a'
	}

	return string(result)
}

// @lc code=end
