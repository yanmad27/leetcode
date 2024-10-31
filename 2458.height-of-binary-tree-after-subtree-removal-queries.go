package main

/*
 * @lc app=leetcode id=2458 lang=golang
 *
 * [2458] Height of Binary Tree After Subtree Removal Queries
 */

// @lc code=start

func getHeight(root *TreeNode) int {
	if root == nil {
		return -1
	}
	return 1 + max(getHeight(root.Left), getHeight(root.Right))
}
func treeQueries(root *TreeNode, queries []int) []int {
	deletedMap := make(map[int]int)
	var dfs func(root *TreeNode, deep int)
	dfs = func(root *TreeNode, deep int) {
		if root == nil {
			return
		}
		if root.Left != nil {
			deletedMap[root.Left.Val] = deep + 1 + getHeight(root.Right)
		}
		if root.Right != nil {
			deletedMap[root.Right.Val] = deep + 1 + getHeight(root.Left)
		}
		dfs(root.Left, deep+1)
		dfs(root.Right, deep+1)
	}
	dfs(root, 0)
	result := []int{}
	for _, query := range queries {
		result = append(result, deletedMap[query])
	}
	return result
}

// @lc code=end
