package main

import (
	"fmt"
	"math"
	"reflect"
)

var passed, total int

func expect(result, expectation any) bool {
	total++
	if !equal(result, expectation) {
		fmt.Printf("FAIL #%d: expected %v (%T), but got %v (%T)\n",
			total, expectation, expectation, result, result)
		return false
	}
	passed++
	return true
}

func report() {
	fmt.Printf("%d/%d passed\n", passed, total)
}

// equal compares numbers by value, so expect(float64(5), 5) passes even though
// the untyped constant arrives as an int. Floats are compared with a small
// tolerance. Everything else falls back to DeepEqual.
func equal(result, expectation any) bool {
	if a, ok := asFloat(result); ok {
		if b, ok := asFloat(expectation); ok {
			return math.Abs(a-b) < 1e-9
		}
	}
	return reflect.DeepEqual(result, expectation)
}

func asFloat(v any) (float64, bool) {
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(rv.Int()), true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(rv.Uint()), true
	case reflect.Float32, reflect.Float64:
		return rv.Float(), true
	}
	return 0, false
}
