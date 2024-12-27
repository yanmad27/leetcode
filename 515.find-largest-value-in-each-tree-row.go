package main

/*
 * @lc app=leetcode id=515 lang=golang
 *
 * [515] Find Largest Value in Each Tree Row
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
func largestValues(root *TreeNode) []int {
	res := []int{}
	var travel func(*TreeNode, int)
	travel = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if level == len(res) {
			res = append(res, node.Val)
		} else {
			if node.Val >= res[level] {
				res[level] = node.Val
			}
		}
		travel(node.Left, level+1)
		travel(node.Right, level+1)
	}
	travel(root, 0)
	return res
}

// @lc code=end
