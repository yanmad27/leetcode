package main

import "fmt"

/*
 * @lc app=leetcode id=951 lang=golang
 *
 * [951] Flip Equivalent Binary Trees
 */

// @lc code=start
func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {

	path1Map := make(map[string]int)
	path2Map := make(map[string]int)

	var getPath func(tmp int, root *TreeNode, curPath string)
	getPath = func(tmp int, root *TreeNode, curPath string) {
		if root == nil {
			if tmp == 1 {
				path1Map[curPath]++
			} else {
				path2Map[curPath]++

			}
			return
		}
		curPath = fmt.Sprintf("%s_%d", curPath, root.Val)
		getPath(tmp, root.Left, curPath)
		getPath(tmp, root.Right, curPath)
	}
	getPath(1, root1, "")
	getPath(2, root2, "")
	for key1, val1 := range path1Map {
		val2, ok := path2Map[key1]
		if val2 != val1 || !ok {
			return false
		}
	}
	for key2, val2 := range path2Map {
		val1, ok := path1Map[key2]
		if val2 != val1 || !ok {
			return false
		}
	}
	return true

}

// @lc code=end
