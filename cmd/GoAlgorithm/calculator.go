package main

type Node struct {
	Val   string
	Left  *Node
	Right *Node
}

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
	tokens := GetTokens(input)
	if len(tokens) > 0 {
		postfixTokens := ParseExpression(tokens)
		if len(postfixTokens) > 0 {
			return makeTreeFromPostfixTokens(postfixTokens)
		}
	}
	return nil
}

func makeTreeFromPostfixTokens(tokens []string) *Node {
	var nodeStack []*Node
	for _, token := range tokens {
		newNode := createCalculateNode(token)
		if IsNumber(token) {
			nodeStack = append(nodeStack, newNode)
		} else {
			last := nodeStack[len(nodeStack)-1]
			secondLast := nodeStack[len(nodeStack)-2]
			newNode.Left = secondLast
			newNode.Right = last
			nodeStack = nodeStack[:len(nodeStack)-2]
			nodeStack = append(nodeStack, newNode)
		}
	}
	return nodeStack[0]
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
