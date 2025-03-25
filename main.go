package main

import "fmt"

func main() {
	fmt.Println(checkValidCuts(5, convertTo2DArray("[[1,0,5,2],[0,2,2,4],[3,2,5,3],[0,4,4,5]]")))
	fmt.Println(checkValidCuts(4, convertTo2DArray("[[0,0,1,1],[2,0,3,4],[0,2,2,3],[3,0,4,3]]")))
}
