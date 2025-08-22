package tree

import (
	"math"
)

type BinaryTree struct {
	Root *BinaryTreeNode
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

func (bt *BinaryTree) Min() (float64, bool) {
	if bt.Root == nil {
		return math.NaN(), false
	}
	return MinHelp(bt.Root), true
}

func MinHelp(node *BinaryTreeNode) float64 {
	result := node.Value
	if node.Left != nil {
		leftMin := MinHelp(node.Left)
		result = math.Min(result, leftMin)
	}
	if node.Right != nil {
		rightMin := MinHelp(node.Right)
		result = math.Min(result, rightMin)
	}
	return result
}

func (bt *BinaryTree) LevelOrder() {
	if bt.Root == nil {
		return
	}

	var q []BinaryTreeNode
	q = append(q, *bt.Root)
	for len(q) > 0 {
		current := q[0]
		q = q[1:]
		println(current.Value)

		if current.Left != nil {
			q = append(q, *current.Left)
		}
		if current.Right != nil {
			q = append(q, *current.Right)
		}
	}
}

func (bt *BinaryTree) FindMaxSpan() int {
	if bt.Root == nil {
		return 0
	}

	result, _ := FindMaxSpanHelp(bt.Root)
	return int(result)
}

func FindMaxSpanHelp(node *BinaryTreeNode) (float64, float64) {
	if node == nil {
		return 0, 0
	}

	leftResult, leftHeight := FindMaxSpanHelp(node.Left)
	rightResult, rightHeight := FindMaxSpanHelp(node.Right)
	height := math.Max(leftHeight, rightHeight) + 1

	crossNodeSpan := leftHeight + rightHeight + 1
	result := math.Max(crossNodeSpan, math.Max(leftResult, rightResult))
	return result, height
}
