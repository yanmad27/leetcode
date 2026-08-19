package main

import (
	"fmt"
	"reflect"
)

var passed, total int

func expect(result, expectation interface{}) {
	total++
	if !reflect.DeepEqual(result, expectation) {
		fmt.Printf("FAIL #%d: expected %v, but got %v\n", total, expectation, result)
		return
	}
	passed++
}

func report() {
	fmt.Printf("%d/%d passed\n", passed, total)
}
