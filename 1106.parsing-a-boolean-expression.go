package main

import (
	"slices"
	"strings"
)

/*
 * @lc app=leetcode id=1106 lang=golang
 *
 * [1106] Parsing A Boolean Expression
 */

// @lc code=start

func operatorNOT(value string) bool {
	if value[0] != 't' && value[0] != 'f' {
		return !parseBoolExpr(value)
	}
	return value != "t"
}
func operatorAND(value string) bool {
	for i := 0; i < len(value); i++ {
		operator := getOperator(value[i])
		if operator != nil {
			operatorValue := getValue(value, i)
			result := operator(operatorValue)
			tmp := "f"
			if result {
				tmp = "t"
			}
			value = value[:i] + tmp + value[3+i+len(operatorValue):]
		}
	}

	split := strings.Split(value, ",")
	return !slices.Contains(split, "f")
}
func operatorOR(value string) bool {
	for i := 0; i < len(value); i++ {
		operator := getOperator(value[i])
		if operator != nil {
			operatorValue := getValue(value, i)
			result := operator(operatorValue)
			tmp := "f"
			if result {
				tmp = "t"
			}
			value = value[:i] + tmp + value[3+i+len(operatorValue):]
		}
	}

	split := strings.Split(value, ",")
	return slices.Contains(split, "t")
}

type operator func(value string) bool

func getOperator(c byte) operator {
	switch c {
	case '!':
		return operatorNOT
	case '&':
		return operatorAND
	case '|':
		return operatorOR
	default:
		return nil
	}
}

func getValue(expression string, operatorIndex int) string {
	openCount := 0
	closeCount := 0
	for i := operatorIndex + 1; i < len(expression); i++ {
		if expression[i] == '(' {
			openCount++
		} else if expression[i] == ')' {
			closeCount++
		}
		if openCount == closeCount {
			return expression[operatorIndex+2 : i]
		}
	}
	panic("Unreachable")
}

func parseBoolExpr(expression string) bool {
	result := false
	operator := getOperator(expression[0])
	if operator != nil {
		value := getValue(expression, 0)
		result = operator(value)

	}
	return result
}

// @lc code=end
