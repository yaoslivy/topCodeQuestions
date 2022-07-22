package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	// return insertIntoBSTRecursion(root, val)

	return insertIntoBSTTraverse(root, val)
}

//迭代方式
func insertIntoBSTTraverse(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	cur := root
	var pre *TreeNode
	for cur != nil {
		pre = cur
		if cur.Val > val {
			cur = cur.Left
		} else if cur.Val < val {
			cur = cur.Right
		}
	}
	if pre.Val > val {
		pre.Left = &TreeNode{Val: val}
	} else {
		pre.Right = &TreeNode{Val: val}
	}
	return root
}

//递归方式
func insertIntoBSTRecursion(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if root.Val > val { //在当前节点的左边
		root.Left = insertIntoBSTRecursion(root.Left, val)
	}
	if root.Val < val { //在当前节点的右边
		root.Right = insertIntoBSTRecursion(root.Right, val)
	}
	return root
}
