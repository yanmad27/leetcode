package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//	func parse1DArr(text string) []int {
//		text = strings.ReplaceAll(text, "[", "")
//		text = strings.ReplaceAll(text, "]", "")
//		arr := strings.Split(text, ",")
//		rs := []int{}
//		for _, _ := range arr {
//			rs = append(rs, 1)
//		}
//		return rs
//	}
func buildTreeNode(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: nums[0],
	}
	queue := make([]*TreeNode, 1)
	queue[0] = root
	i := 1
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if i < len(nums) && nums[i] != -1 {
			node.Left = &TreeNode{
				Val: nums[i],
			}
			queue = append(queue, node.Left)
		}
		i++
		if i < len(nums) && nums[i] != -1 {
			node.Right = &TreeNode{
				Val: nums[i],
			}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}
