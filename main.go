package main

import (
	"fmt"
	"reflect"
)

func expect(result, expectation interface{}) {
	if !reflect.DeepEqual(result, expectation) {
		panic(fmt.Sprintf("Expected %v, but got %v", expectation, result))
	}

}
func main() {
	expect(compress([]byte{'a', 'b', 'c', 'c', 'c', 'c', 'c', 'c', 'd', 'd'}), "asdf")
}
