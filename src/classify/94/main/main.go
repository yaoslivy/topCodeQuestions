package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res []int

func inorderTraversal(root *TreeNode) []int {
	res = make([]int, 0)
	inOrder(root)
	// inTraverse(root)
	return res
}

//中序递归
func inOrder(root *TreeNode) {
	if root == nil {
		return
	}
	inOrder(root.Left)
	res = append(res, root.Val)
	inOrder(root.Right)
}

//中序迭代
func inTraverse(root *TreeNode) {
	S := make([]*TreeNode, 0)
	cur := root
	for cur != nil || len(S) != 0 {
		if cur != nil {
			S = append(S, cur)
			cur = cur.Left
		} else {
			cur = S[len(S)-1]
			S = S[:len(S)-1]
			res = append(res, cur.Val)
			cur = cur.Right
		}
	}
}
