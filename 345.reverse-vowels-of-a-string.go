package main

/*
 * @lc app=leetcode id=345 lang=golang
 *
 * [345] Reverse Vowels of a String
 */

// @lc code=start
func reverseVowels(s string) string {
	n := len(s)
	arr := []byte(s)
	l, r := 0, n-1
	vowelMap := map[byte]bool{
		'a': true,
		'i': true,
		'u': true,
		'e': true,
		'o': true,
		'A': true,
		'I': true,
		'U': true,
		'E': true,
		'O': true,
	}
	for l < r {

		for l < n && !vowelMap[s[l]] {
			l++
		}
		for r >= 0 && !vowelMap[s[r]] {
			r--
		}
		if l >= r {
			break
		}
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
	return string(arr)
}

// @lc code=end
