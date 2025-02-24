package main

import "fmt"

func main() {
	fmt.Println(mostProfitablePath(
		convertTo2DArray("[[0,2],[1,4],[1,6],[2,4],[3,6],[3,7],[5,7]]"),
		4,
		convertTo1DArray("[-6896,-1216,-1208,-1108,1606,-7704,-9212,-8258]"),
	))
}
