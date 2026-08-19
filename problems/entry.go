package problems

import "fmt"

// current is set by the file ./solve generates for whichever problem you are
// working on, and deletes once the run is over.
var current func()

// Main runs the selected problem and reports how many cases passed.
func Main() {
	defer report()

	if current == nil {
		fmt.Println("no problem selected: run ./solve")
		return
	}
	current()
}
