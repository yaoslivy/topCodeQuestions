package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res []int

func postorderTraversal(root *TreeNode) []int {
	res = make([]int, 0)
	// postOrder(root)
	postTarverse(root)
	return res
}

//后序递归
func postOrder(root *TreeNode) {
	if root == nil {
		return
	}
	postOrder(root.Left)
	postOrder(root.Right)
	res = append(res, root.Val)
}

//后序迭代
func postTarverse(root *TreeNode) {
	S := make([]*TreeNode, 0)
	cur := root
	var pre *TreeNode //记录当前前一个访问的节点
	for cur != nil || len(S) != 0 {
		if cur != nil {
			S = append(S, cur)
			cur = cur.Left
		} else {
			cur = S[len(S)-1]
			//当前节点的右子树不为nil，并且没有访问过
			if cur.Right != nil && cur.Right != pre {
				cur = cur.Right
			} else {
				S = S[:len(S)-1]
				res = append(res, cur.Val)
				pre = cur
				cur = nil
			}
		}
	}
}
