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
	var pre *TreeNode //记录cur的父节点
	//找到合适的空位置插入
	for cur != nil {
		pre = cur
		if cur.Val >= val {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	if pre.Val >= val { //插入左边
		pre.Left = &TreeNode{Val: val}
	} else {
		pre.Right = &TreeNode{Val: val}
	}
	return root
}

//递归
func insertIntoBSTRecursion(root *TreeNode, val int) *TreeNode {
	//递归找到一个空位置插入
	if root == nil {
		return &TreeNode{Val: val}
	}

	if root.Val >= val { //插入左边
		root.Left = insertIntoBSTRecursion(root.Left, val)
	} else {
		root.Right = insertIntoBSTRecursion(root.Right, val)
	}
	return root
}
