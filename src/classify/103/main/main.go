package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	Q := make([]*TreeNode, 0)
	p := root
	Q = append(Q, p)
	res := make([][]int, 0)
	count := 1
	for len(Q) != 0 {
		size := len(Q)
		oneLevel := make([]int, 0)
		for ; size > 0; size-- { //逐层添加
			p = Q[0]
			if count%2 == 0 { //偶数行逆序
				oneLevel = append([]int{p.Val}, oneLevel...)
			} else {
				oneLevel = append(oneLevel, p.Val)
			}
			Q = Q[1:]
			if p.Left != nil {
				Q = append(Q, p.Left)
			}
			if p.Right != nil {
				Q = append(Q, p.Right)
			}
		}
		res = append(res, oneLevel)
		count++
	}
	return res
}
