package main

import "strconv"

func calculate(tokens []string) int {
	result := 0
	if len(tokens) == 0 {
		return result
	}

	var stack []string
	for _, token := range tokens {
		if IsNumber(token) {
			stack = append(stack, token)
		} else {
			last := stack[len(stack)-1]
			secondLast := stack[len(stack)-2]
			lastInt, _ := strconv.Atoi(last)
			secondLastInt, _ := strconv.Atoi(secondLast)
			midResult := 0
			if token == "+" {
				midResult = secondLastInt + lastInt
			} else if token == "-" {
				midResult = secondLastInt - lastInt
			} else if token == "*" {
				midResult = secondLastInt * lastInt
			} else if token == "/" {
				midResult = secondLastInt / lastInt
			}
			stack = stack[:len(stack)-2]
			stack = append(stack, strconv.Itoa(midResult))
		}
	}
	result, _ = strconv.Atoi(stack[0])
	return result
}

func ParseExpression(tokens []string) []string {
	var operators []string
	var data []string
	for _, token := range tokens {
		if IsNumber(token) {
			data = append(data, token)
		} else {
			if IsOperator(token) {
				if len(operators) > 0 {
					lastOperator := operators[len(operators)-1]
					if isOperatorHigher(token, lastOperator) {
						operators = append(operators, token)
					} else {
						for _, operator := range operators {
							if !isOperatorHigher(operator, token) || operator == "(" {
								data = append(data, operator)
								operators = operators[:len(operators)-2]
							}
						}
						data = append(data, token)
					}
				} else {
					operators = append(operators, token)
				}
			} else if token == "(" {
				operators = append(operators, token)
			} else if token == ")" {
				for _, operator := range operators {
					if operator != "(" {
						data = append(data, operator)
						operators = operators[:len(operators)-2]
					}
				}
				operators = operators[:len(operators)-2]
			}
		}
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

func IsOperator(source string) bool {
	result := false
	if source == "+" || source == "-" || source == "*" || source == "/" {
		result = true
	}
	return result
}

func IsOperatorHigh(left string, right string) bool {
	result := false
	if left == "*" || left == "/" {
		if right == "+" || right == "-" {
			result = true
		}
	}
	return result
}
