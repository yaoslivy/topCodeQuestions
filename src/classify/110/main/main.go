package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	//判断每一个子节点为根节点的左右子树的高度差
	leftD := getDepth(root.Left)
	rightD := getDepth(root.Right)
	return math.Abs(float64(leftD-rightD)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

//递归找到一个子树的高度
func getDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftD := getDepth(root.Left)
	rightD := getDepth(root.Right)

	return max(leftD, rightD) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
