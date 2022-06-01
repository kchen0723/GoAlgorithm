package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var inorderMap map[int]int

func buildTree(inorder []int, postorder []int) *TreeNode {
	if inorder == nil && postorder == nil && len(inorder) == 0 && len(inorder) != len(postorder) {
		return nil
	}
	inorderMap = make(map[int]int)
	buildInorderMap(inorder)

	// node := new(TreeNode)
	// node.Val = postorder[len(postorder) - 1]
	// return node

	var result = buildTreeHelper(inorder, 0, len(inorder)-1, postorder, 0, len(postorder)-1)
	return result
}

func buildTreeHelper(inorder []int, inorderStart int, inorderEnd int,
	postorder []int, postorderStart int, postorderEnd int) *TreeNode {
	if inorderStart < 0 {
		return nil
	}
	node := new(TreeNode)
	node.Val = postorder[postorderEnd]

	return node
}

func buildInorderMap(inorder []int) {
	for i, item := range inorder {
		_, isExist := inorderMap[item]
		if !isExist {
			inorderMap[item] = i
		}
	}
}
