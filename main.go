package main

import "fmt"

// current is set by the file ./solve generates for whichever problem you are
// working on, and deletes once the run is over. This file stays as it is.
var current func()

func main() {
	defer report()

	if current == nil {
		fmt.Println("no problem selected: run ./solve")
		return
	}
	current()
}
