package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}
	Q := make([]*TreeNode, 0)
	p := root
	Q = append(Q, p)
	res := make([]float64, 0)
	for len(Q) != 0 {
		//当层节点数量
		size := len(Q)
		sum := 0.0 //每一层的求和
		n := float64(size)
		for size != 0 {
			p = Q[0]
			Q = Q[1:]
			if p.Left != nil {
				Q = append(Q, p.Left)
			}
			if p.Right != nil {
				Q = append(Q, p.Right)
			}
			sum += float64(p.Val)
			size--
		}
		res = append(res, sum/n)
	}
	return res
}
