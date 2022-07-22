package main

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}
	Q := make([]*Node, 0)
	p := root
	Q = append(Q, p)
	res := make([][]int, 0)
	for len(Q) != 0 {
		//当层节点个数
		size := len(Q)
		oneRes := make([]int, 0)
		for size != 0 {
			p = Q[0]
			Q = Q[1:]
			if len(p.Children) != 0 {
				cSize := len(p.Children) // 每个节点的子节点集合个数
				for i := 0; i < cSize; i++ {
					Q = append(Q, p.Children[i])
				}
			}
			oneRes = append(oneRes, p.Val)
			size--
		}
		res = append(res, oneRes)
	}
	return res
}
