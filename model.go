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

// func parse1DArr(text string) []int {
// 	text = strings.ReplaceAll(text, "[", "")
// 	text = strings.ReplaceAll(text, "]", "")
// 	arr := strings.Split(text, ",")
// 	rs := []int{}
// 	for _, _ := range arr {
// 		rs = append(rs, 1)
// 	}
// 	return rs
// }
