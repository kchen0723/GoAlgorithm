package tree

import (
	"fmt"
	"testing"
)

func TestMin(t *testing.T) {
	tree := createTestingBinaryTree()
	min, _ := tree.Min()
	fmt.Println(min)
}

func TestLevelOrder(t *testing.T) {
	tree := createTestingBinaryTree()
	tree.LevelOrder()
}

func TestFindMaxSpan(t *testing.T) {
	tree := createTestingBinaryTree()
	maxSpan := tree.FindMaxSpan()
	fmt.Println(maxSpan)
}

func TestInOrderIteration(t *testing.T) {
	tree := createTestingBinaryTree()
	tree.InOrderIteration()
}

func createTestingBinaryTree() *BinaryTree {
	result := NewBinaryTree()
	result.Root = &BinaryTreeNode{
		Value: 8,
		Left: &BinaryTreeNode{
			Value: 7,
			Left: &BinaryTreeNode{
				Value: 6,
				Left: &BinaryTreeNode{
					Value: 9,
				},
			},
			Right: &BinaryTreeNode{
				Value: 5,
				Left: &BinaryTreeNode{
					Value: 2,
				},
				Right: &BinaryTreeNode{
					Value: 4,
				},
			},
		},
		Right: &BinaryTreeNode{
			Value: 3,
		},
	}
	return result
}
