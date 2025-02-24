package main

import "fmt"

func main() {
	fmt.Println(mostProfitablePath(
		convertTo2DArray("[[0,1],[1,2],[1,3],[3,4]]"),
		3,
		convertTo1DArray("[-2,4,2,-4,6]"),
	))
}
