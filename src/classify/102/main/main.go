package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	p := root
	Q := make([]*TreeNode, 0)
	Q = append(Q, p)
	res := make([][]int, 0)
	for len(Q) != 0 {
		size := len(Q) //每一层的数量
		oneLevel := make([]int, 0)
		for ; size > 0; size-- {
			p = Q[0]
			Q = Q[1:]
			oneLevel = append(oneLevel, p.Val)
			if p.Left != nil {
				Q = append(Q, p.Left)
			}
			if p.Right != nil {
				Q = append(Q, p.Right)
			}
		}
		res = append(res, oneLevel)
	}
	return res
}
