package main

import "fmt"

func main() {
	fmt.Println(countDays(10, convertTo2DArray("[[5,7],[1,3],[9,10]]")))
	fmt.Println(countDays(5, convertTo2DArray("[[2,4],[1,3]]")))
	fmt.Println(countDays(6, convertTo2DArray("[[1,6]")))
	fmt.Println(countDays(10, convertTo2DArray("[[1,4],[2,4],[5,6],[8,9],[9,10]]")))
	fmt.Println(countDays(8, convertTo2DArray("[[3,4],[4,8],[2,5],[3,8]]")))
}
