package main

import "fmt"

func main() {
	// fmt.Println(compress([]byte{'a', 'a', 'a', 'b', 'b', 'a', 'a'}))
	//["a","b","c","c","c","c","c","c"]
	fmt.Println(compress([]byte{'a', 'b', 'c', 'c', 'c', 'c', 'c', 'c', 'd', 'd'}))
}
