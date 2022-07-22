package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	Q := make([]*TreeNode, 0)
	p := root
	Q = append(Q, p)
	res := make([][]int, 0)
	for len(Q) != 0 {
		size := len(Q) //当前层的个数
		oneRes := make([]int, 0)
		for size != 0 { //对当前层的节点做遍历
			p = Q[0]
			Q = Q[1:]
			if p.Left != nil {
				Q = append(Q, p.Left)
			}
			if p.Right != nil {
				Q = append(Q, p.Right)
			}
			size--
			oneRes = append(oneRes, p.Val)
		}
		res = append([][]int{oneRes}, res...) //插入结果集的首部
	}
	return res
}
