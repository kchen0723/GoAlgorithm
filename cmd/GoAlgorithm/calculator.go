package main

type Node struct {
	Val   string
	Left  *Node
	Right *Node
}

var calculateTree *Node
var index int
var dataStack []string
var operatorStack []string

func sealize(node *Node) string {
	if node == nil {
		return ""
	}
	sealizeString := sealizeHelper(node)
	return sealizeString
}

func sealizeHelper(node *Node) string {
	if node == nil {
		return ""
	}
	left := sealizeHelper(node.Left)
	isLeftAddBracket := shouldAddBracket(node, node.Left)
	right := sealizeHelper(node.Right)
	isRightAddBracket := shouldAddBracket(node, node.Right)
	result := ""
	if isLeftAddBracket {
		result += "(" + left + ")"
	} else {
		result += left
	}
	result += node.Val
	if isRightAddBracket {
		result += "(" + right + ")"
	} else {
		result += right
	}
	return result
}

func desealize(input string) *Node {
	if len(input) == 0 {
		return nil
	}
	index = 0
	return nil
}

func ToPostfixNotation(input string) {
	index = 0
	token := getToken(input)
	for len(token) > 0 {
		if !isSpecial(token) {
			dataStack = append(dataStack, token)
		} else {
			if token == "(" {
				operatorStack = append(operatorStack, token)
			} else if token == ")" {
				for len(operatorStack) > 0 {
					lastOperator := operatorStack[len(operatorStack)-1]
					operatorStack = operatorStack[:len(operatorStack)-1]
					if lastOperator != "(" {
						dataStack = append(dataStack, lastOperator)
					} else {
						break
					}
				}
			} else {
				if len(operatorStack) > 0 {
					lastOperator := operatorStack[len(operatorStack)-1]
					if lastOperator == "(" || lastOperator == ")" {
						operatorStack = append(operatorStack, token)
					} else if (token == "*" || token == "/") && (lastOperator == "+" || lastOperator == "-") {
						operatorStack = append(operatorStack, token)
					} else {
						dataStack = append(dataStack, lastOperator)
						operatorStack[len(operatorStack)-1] = token
					}
				} else {
					operatorStack = append(operatorStack, token)
				}
			}
		}
		token = getToken(input)
	}
	for i := len(operatorStack) - 1; i >= 0; i-- {
		dataStack = append(dataStack, operatorStack[i])
	}
	operatorStack = []string{}
}

func getToken(input string) string {
	if len(input) == 0 {
		return ""
	}
	if index < 0 || index >= len(input) {
		return ""
	}
	result := ""
	single := string(input[index])
	isOperator := isSpecial(single)
	if isOperator {
		result = single
		index++
	} else {
		for i := 0; i < len(input)-index; i++ {
			currentSingle := string(input[index+i])
			if isSpecial(currentSingle) {
				result = input[index : index+i]
				index += i
				break
			} else if i == len(input)-index-1 {
				result = input[index : index+i+1]
				index += i
				break
			}
		}
	}
	return result
}

func isSpecial(single string) bool {
	if single == "(" || single == ")" || single == "+" || single == "-" || single == "*" || single == "/" {
		return true
	}
	return false
}

func shouldAddBracket(left *Node, right *Node) bool {
	result := false
	if left != nil && right != nil {
		result = isOperatorHigher(left.Val, right.Val)
	}
	return result
}

func isOperatorHigher(left string, right string) bool {
	result := false
	if left == "*" || left == "/" {
		if right == "+" || right == "-" {
			result = true
		}
	}
	return result
}

func createCalculateTree() *Node {
	eight := createCalculateNode("8")
	four := createCalculateNode("4")

	first := createCalculateNode("1")
	divid := createCalculateNode("/")
	divid.Left = eight
	divid.Right = four
	add := createCalculateNode("+")
	add.Left = first
	add.Right = divid

	fifth := createCalculateNode("5")
	three := createCalculateNode("3")
	minus := createCalculateNode("-")
	minus.Left = fifth
	minus.Right = three

	mul := createCalculateNode("*")
	mul.Left = add
	mul.Right = minus
	return mul
}

func createCalculateNode(nodeValue string) *Node {
	node := new(Node)
	node.Val = nodeValue
	return node
}
