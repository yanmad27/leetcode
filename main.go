package main

import "fmt"

func main() {
	list := []int{5, 8, 9, 2, 1, 3, 7, 4, 6}
	fmt.Println(treeQueries(buildTreeNode(list), []int{3, 2, 4, 8}))
}
