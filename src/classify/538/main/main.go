package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	pre = nil
	convertBSTRecursion(root)
	return root
}

// 通过递归从右中左遍历，然后将当前节点值+前一个访问节点值
var pre *TreeNode //记录前一个访问节点
func convertBSTRecursion(root *TreeNode) {
	if root == nil {
		return
	}
	convertBSTRecursion(root.Right)
	if pre != nil {
		root.Val += pre.Val
	}
	pre = root
	convertBSTRecursion(root.Left)
}
