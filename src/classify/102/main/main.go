package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	// 使用队列记录每一层访问过的节点
	if root == nil {
		return [][]int{}
	}
	Q := make([]*TreeNode, 0)
	cur := root
	Q = append(Q, cur)
	res := make([][]int, 0)
	for len(Q) != 0 {
		// 根据当前层的节点个数来添加下一层节点
		size := len(Q)
		oneRes := make([]int, 0)
		for i := 0; i < size; i++ {
			cur = Q[0]
			Q = Q[1:]
			oneRes = append(oneRes, cur.Val)
			if cur.Left != nil {
				Q = append(Q, cur.Left)
			}
			if cur.Right != nil {
				Q = append(Q, cur.Right)
			}
		}
		res = append(res, oneRes)
	}
	return res
}
