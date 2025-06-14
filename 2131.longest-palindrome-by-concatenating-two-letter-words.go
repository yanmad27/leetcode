package main

import (
	"strconv"
	"strings"
)

func CreatePhoneNumber(numbers [10]uint) string {

	builder := strings.Builder{}
	// "(123) 456-7890"
	builder.WriteString("(")
	builder.WriteString(strconv.Itoa(int(numbers[0])))
	builder.WriteString(strconv.Itoa(int(numbers[1])))
	builder.WriteString(strconv.Itoa(int(numbers[2])))
	builder.WriteString(") ")
	builder.WriteString(strconv.Itoa(int(numbers[3])))
	builder.WriteString(strconv.Itoa(int(numbers[4])))
	builder.WriteString(strconv.Itoa(int(numbers[5])))
	builder.WriteString("-")
	builder.WriteString(strconv.Itoa(int(numbers[6])))
	builder.WriteString(strconv.Itoa(int(numbers[7])))
	builder.WriteString(strconv.Itoa(int(numbers[8])))
	builder.WriteString(strconv.Itoa(int(numbers[9])))
	return builder.String()
}
