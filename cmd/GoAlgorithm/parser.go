package main

import (
	"strconv"
)

func calculate(tokens []string) (float64, bool) {
	result := float64(0)
	if len(tokens) == 0 {
		return result, false
	}

	var stack []string
	for _, token := range tokens {
		if IsNumber(token) {
			stack = append(stack, token)
		} else {
			last := stack[len(stack)-1]
			secondLast := stack[len(stack)-2]
			lastFloat64, _ := strconv.ParseFloat(last, 64)
			secondLastFloat64, _ := strconv.ParseFloat(secondLast, 64)
			midResult := float64(0)
			if token == "+" {
				midResult = secondLastFloat64 + lastFloat64
			} else if token == "-" {
				midResult = secondLastFloat64 - lastFloat64
			} else if token == "*" {
				midResult = secondLastFloat64 * lastFloat64
			} else if token == "/" {
				if lastFloat64 == 0 {
					return float64(0), false
				}
				midResult = secondLastFloat64 / lastFloat64
			}
			stack = stack[:len(stack)-2]
			stack = append(stack, strconv.FormatFloat(midResult, 'f', -1, 64))
		}
	}
	result, _ = strconv.ParseFloat(stack[0], 64)
	return result, true
}

func ParseExpression(tokens []string) []string {
	var operators []string
	var data []string
	for _, token := range tokens {
		if IsNumber(token) {
			data = append(data, token)
		} else if token == "(" || len(operators) == 0 {
			operators = append(operators, token)
		} else if token == ")" {
			for i := len(operators) - 1; i >= 0; i-- {
				if operators[i] == "(" {
					operators = operators[:len(operators)-1]
					break
				} else {
					data = append(data, operators[i])
					operators = operators[:len(operators)-1]
				}
			}
		} else {
			for i := len(operators) - 1; i >= 0; i-- {
				lastOperator := operators[i]
				if lastOperator == "(" ||
					((token == "*" || token == "/") && (lastOperator == "+" || lastOperator == "-")) {
					break
				}
				data = append(data, lastOperator)
				operators = operators[:len(operators)-1]
			}
			operators = append(operators, token)
		}
	}
	for i := len(operators) - 1; i >= 0; i-- {
		data = append(data, operators[i])
	}
	return data
}

func IsNumber(source string) bool {
	result := true
	if source == ")" || source == "(" || source == "+" || source == "-" || source == "*" || source == "/" {
		result = false
	}
	return result
}
