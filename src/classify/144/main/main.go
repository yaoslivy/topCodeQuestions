package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res []int

func preorderTraversal(root *TreeNode) []int {
	res = make([]int, 0)
	// preOrder(root)
	preTraverse(root)
	return res
}

//先序递归
func preOrder(root *TreeNode) {
	if root == nil {
		return
	}
	res = append(res, root.Val)
	preOrder(root.Left)
	preOrder(root.Right)
}

//先序迭代
func preTraverse(root *TreeNode) {
	S := make([]*TreeNode, 0)
	cur := root
	for cur != nil || len(S) != 0 {
		if cur != nil {
			S = append(S, cur)
			res = append(res, cur.Val)
			cur = cur.Left
		} else {
			cur = S[len(S)-1]
			S = S[:len(S)-1]
			cur = cur.Right
		}
	}
}
