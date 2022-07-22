package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	Q := make([]*TreeNode, 0)
	p := root
	Q = append(Q, p)
	res := make([]int, 0)
	for len(Q) != 0 {
		size := len(Q) //当前层的节点个数
		for size != 0 {
			//找到当前层的最后一个元素然后添加入结果集合中
			p = Q[0]
			Q = Q[1:]
			if p.Left != nil {
				Q = append(Q, p.Left)
			}
			if p.Right != nil {
				Q = append(Q, p.Right)
			}
			if size == 1 {
				res = append(res, p.Val)
			}
			size--
		}
	}
	return res
}
