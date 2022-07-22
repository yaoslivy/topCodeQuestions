package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	//根节点不动，交换左右子树
	// head := invert(root)
	// return head

	// invert2(root)
	// return root

	// head := invert3(root)
	// return head

	head := invert4(root)
	return head
}

//后序遍历 迭代
func invert4(root *TreeNode) *TreeNode {
	S := make([]*TreeNode, 0)
	cur := root
	var pre *TreeNode
	for cur != nil || len(S) != 0 {
		if cur != nil {
			S = append(S, cur)
			cur = cur.Left
		} else {
			cur = S[len(S)-1]
			//观察当前栈顶右孩子
			if cur.Right != nil && cur.Right != pre {
				cur = cur.Right
			} else {
				S = S[:len(S)-1]
				pre = cur
				cur.Left, cur.Right = cur.Right, cur.Left
				cur = nil
			}
		}
	}
	return root
}

// 先序遍历 迭代
func invert3(root *TreeNode) *TreeNode {
	S := make([]*TreeNode, 0)
	cur := root
	for cur != nil || len(S) != 0 {
		if cur != nil {
			cur.Left, cur.Right = cur.Right, cur.Left
			S = append(S, cur)
			cur = cur.Left
		} else {
			cur = S[len(S)-1]
			S = S[:len(S)-1]
			cur = cur.Right
		}
	}
	return root
}

//从上往下交换 先序遍历
func invert2(root *TreeNode) {
	if root == nil {
		return
	}
	root.Left, root.Right = root.Right, root.Left
	invert2(root.Left)
	invert2(root.Right)
}

//从下往上交换 后序遍历
func invert(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invert(root.Left)
	right := invert(root.Right)
	root.Left = right
	root.Right = left
	return root
}
