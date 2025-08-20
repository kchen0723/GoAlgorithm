package dp

import "strconv"

func Calcualte(str string) float64 {
	if len(str) == 0 {
		return 0
	}

	runes := []rune(str)
	return calcualteHelp(&runes)
}

func calcualteHelp(runes *[]rune) float64 {
	var stack []float64
	var tokens []rune
	var num float64
	sign := '+'
	for len(*runes) > 0 {
		token := (*runes)[0]
		*runes = (*runes)[1:]
		if (token >= '0' && token <= '9') || token == '.' {
			tokens = append(tokens, token)
		}
		if token == '(' || token == ')' || token == '+' || token == '-' || token == '*' || token == '/' || len(*runes) == 0 {
			if len(tokens) > 0 {
				num, _ = strconv.ParseFloat(string(tokens), 64)
				tokens = tokens[:0]
			}
			if token == '(' {
				num = calcualteHelp(runes)
			}
			switch sign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, 0-num)
			case '*':
				stack[len(stack)-1] = stack[len(stack)-1] * num
			case '/':
				stack[len(stack)-1] = stack[len(stack)-1] / num
			}
			if token == ')' {
				break
			}
			sign = token
		}
	}
	var result float64
	for _, item := range stack {
		result += item
	}
	return result
}
