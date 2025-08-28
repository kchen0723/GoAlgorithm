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
	return minHelp(bt.Root), true
}

func minHelp(node *BinaryTreeNode) float64 {
	result := node.Value
	if node.Left != nil {
		leftMin := minHelp(node.Left)
		result = math.Min(result, leftMin)
	}
	if node.Right != nil {
		rightMin := minHelp(node.Right)
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

	result, _ := findMaxSpanHelp(bt.Root)
	return int(result)
}

func findMaxSpanHelp(node *BinaryTreeNode) (float64, float64) {
	if node == nil {
		return 0, 0
	}

	leftResult, leftHeight := findMaxSpanHelp(node.Left)
	rightResult, rightHeight := findMaxSpanHelp(node.Right)
	height := math.Max(leftHeight, rightHeight) + 1

	crossNodeSpan := leftHeight + rightHeight + 1
	result := math.Max(crossNodeSpan, math.Max(leftResult, rightResult))
	return result, height
}

func (bt *BinaryTree) IsSymmetric() bool {
	if bt.Root == nil {
		return true
	}
	return isSymmetricHelp(bt.Root.Left, bt.Root.Right)
}

func isSymmetricHelp(left *BinaryTreeNode, right *BinaryTreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Value != right.Value {
		return false
	}

	leftResult := isSymmetricHelp(left.Left, right.Right)
	rightResult := isSymmetricHelp(left.Right, right.Left)
	return leftResult && rightResult
}

func (bt *BinaryTree) IsBalanced() bool {
	if bt.Root == nil {
		return true
	}
	result, _ := isBalancedHelp(bt.Root)
	return result
}

func isBalancedHelp(node *BinaryTreeNode) (bool, float64) {
	if node == nil {
		return true, 0
	}
	leftResult, leftHeight := isBalancedHelp(node.Left)
	if !leftResult {
		return false, -1
	}
	rightResult, rightHeight := isBalancedHelp(node.Right)
	if !rightResult {
		return false, -1
	}

	height := math.Max(leftHeight, rightHeight) + 1
	if math.Abs(leftHeight-rightHeight) <= 1 {
		return true, height
	}
	return false, -1
}

func (bt *BinaryTree) IsBst() bool {
	if bt.Root == nil {
		return true
	}
	return isBstHelp(bt.Root, nil, nil)
}

func isBstHelp(current *BinaryTreeNode, min *BinaryTreeNode, max *BinaryTreeNode) bool {
	if current == nil {
		return true
	}
	if min != nil && current.Value <= min.Value {
		return false
	}
	if max != nil && current.Value >= max.Value {
		return false
	}

	leftResult := isBstHelp(current.Left, min, current)
	rightResult := isBstHelp(current.Right, current, max)
	return leftResult && rightResult
}

func (bt *BinaryTree) GetLowestCommonAncestor(p *BinaryTreeNode, q *BinaryTreeNode) *BinaryTreeNode {
	if bt.Root == nil {
		return nil
	}
	return getLowestCommonAncestorHelp(bt.Root, p, q)
}

func getLowestCommonAncestorHelp(current *BinaryTreeNode, p *BinaryTreeNode, q *BinaryTreeNode) *BinaryTreeNode {
	if current == nil {
		return nil
	}
	if current.Value == p.Value || current.Value == q.Value {
		return current
	}

	left := getLowestCommonAncestorHelp(current.Left, p, q)
	right := getLowestCommonAncestorHelp(current.Right, p, q)
	if left != nil && right != nil {
		return current
	}
	if left != nil {
		return left
	}
	if right != nil {
		return right
	}
	return nil
}
