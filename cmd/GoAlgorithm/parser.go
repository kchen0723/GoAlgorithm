package main

import (
	"strconv"
)

func calculate(tokens []string) (float64, bool) {
	result := float64(0)
	if len(tokens) == 0 {
		return result, false
	}

	var stack []float64
	for _, token := range tokens {
		if IsNumber(token) {
			tokenFloat, _ := strconv.ParseFloat(token, 64)
			stack = append(stack, tokenFloat)
		} else {
			last := stack[len(stack)-1]
			secondLast := stack[len(stack)-2]
			midResult := float64(0)
			if token == "+" {
				midResult = secondLast + last
			} else if token == "-" {
				midResult = secondLast - last
			} else if token == "*" {
				midResult = secondLast * last
			} else if token == "/" {
				if last == 0 {
					return float64(0), false
				}
				midResult = secondLast / last
			}
			stack = stack[:len(stack)-2]
			stack = append(stack, midResult)
		}
	}
	return stack[0], true
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
