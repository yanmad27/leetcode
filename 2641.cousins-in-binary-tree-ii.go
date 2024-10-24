package main

/*
 * @lc app=leetcode id=2641 lang=golang
 *
 * [2641] Cousins in Binary Tree II
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func replaceValueInTree(root *TreeNode) *TreeNode {
	levelSumMap := make(map[int]int)
	var calculateLevelSum func(root *TreeNode, level int)
	calculateLevelSum = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		levelSumMap[level] += root.Val
		calculateLevelSum(root.Left, level+1)
		calculateLevelSum(root.Right, level+1)
	}
	calculateLevelSum(root, 0)

	var replace func(root *TreeNode, level, sum int)
	replace = func(root *TreeNode, level, sum int) {
		if root == nil {
			return
		}
		root.Val = sum
		curSum := levelSumMap[level+1]
		if root.Left != nil {
			curSum -= root.Left.Val
		}
		if root.Right != nil {
			curSum -= root.Right.Val
		}

		replace(root.Left, level+1, curSum)
		replace(root.Right, level+1, curSum)
	}
	replace(root, 0, 0)
	return root
}

// @lc code=end
