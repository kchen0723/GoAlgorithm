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
		} else if token == "(" {
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
			if len(operators) > 0 {
				for i := len(operators) - 1; i >= 0; i-- {
					lastOperator := operators[len(operators)-1]
					if lastOperator == "(" {
						break
					} else {
						if IsOperatorNotLow(token, lastOperator) {
							data = append(data, lastOperator)
							operators = operators[:len(operators)-1]
						} else {
							break
						}
					}
				}
				operators = append(operators, token)
			} else {
				operators = append(operators, token)
			}
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

func IsOperator(source string) bool {
	result := false
	if source == "+" || source == "-" || source == "*" || source == "/" {
		result = true
	}
	return result
}

func IsOperatorNotLow(token string, lastOperator string) bool {
	result := true
	if lastOperator == "+" || lastOperator == "-" {
		if token == "*" || token == "/" {
			result = false
		}
	}
	return result
}
