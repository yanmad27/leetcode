package main

/*
 * @lc app=leetcode id=2182 lang=golang
 *
 * [2182] Construct String With Repeat Limit
 */

// @lc code=start
func repeatLimitedString(s string, repeatLimit int) string {
	countMap := make(map[rune]int)
	for _, c := range s {
		countMap[c]++
	}

	result := []rune{}
	for i := 25; i >= 0; i-- {
		c := rune(i) + 'a'
		if countMap[c] == 0 {
			continue
		}
		count := 0
		for countMap[c] > 0 {
			result = append(result, c)
			countMap[c]--
			count++
			if countMap[c] > 0 && count%repeatLimit == 0 {
				j := i - 1
				for ; j >= 0; j-- {
					if countMap[rune(j)+'a'] > 0 {
						break
					}
				}
				if j != -1 {
					result = append(result, rune(j)+'a')
					countMap[rune(j)+'a']--
				} else {

					return string(result)
				}
			}
		}
	}
	return string(result)
}

// @lc code=end
