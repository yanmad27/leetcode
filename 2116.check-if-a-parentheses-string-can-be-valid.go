package main

/*
 * @lc app=leetcode id=2116 lang=golang
 *
 * [2116] Check if a Parentheses String Can Be Valid
 */

// @lc code=start
func canBeValid(s string, locked string) bool {
	if len(s)%2 == 1 {
		return false
	}
	unlockStack, openStack := []int{}, []int{}
	for i := 0; i < len(s); i++ {
		if locked[i] == '0' {
			unlockStack = append(unlockStack, i)
		} else if s[i] == '(' {
			openStack = append(openStack, i)
		} else if s[i] == ')' {
			if len(openStack) != 0 {
				openStack = openStack[:len(openStack)-1]
			} else if len(unlockStack) != 0 {
				unlockStack = unlockStack[:len(unlockStack)-1]
			} else {
				return false
			}
		}
	}

	for len(unlockStack) != 0 && len(openStack) != 0 &&
		openStack[len(openStack)-1] < unlockStack[len(unlockStack)-1] {
		openStack = openStack[:len(openStack)-1]
		unlockStack = unlockStack[:len(unlockStack)-1]
	}

	return len(openStack) == 0
}

// @lc code=end
